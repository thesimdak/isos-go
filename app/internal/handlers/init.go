package handlers

import (
	"database/sql"
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/thesimdak/goisos/internal/handlers/upload"
	"github.com/thesimdak/goisos/internal/repository"
	"github.com/thesimdak/goisos/internal/repository/category"
	"github.com/thesimdak/goisos/internal/repository/competition"
	"github.com/thesimdak/goisos/internal/repository/participation"
	"github.com/thesimdak/goisos/internal/repository/result"
	"github.com/thesimdak/goisos/internal/repository/ropeclimber"
	timeRepo "github.com/thesimdak/goisos/internal/repository/time"
	competitionService "github.com/thesimdak/goisos/internal/services/competition"
	resultService "github.com/thesimdak/goisos/internal/services/result"
	ropeClimberService "github.com/thesimdak/goisos/internal/services/ropeclimber"
)

func Initialize(db *sql.DB, staticFS embed.FS) {
	// Attach the Logger and Recovery middleware manually
	// Default to release mode
	gin.SetMode(gin.ReleaseMode)

	// Allow environment variable to override
	if os.Getenv("GIN_MODE") == "debug" {
		gin.SetMode(gin.DebugMode)
	}
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	// Load templates
	router.SetHTMLTemplate(parseTemplates(staticFS))
	router.SetTrustedProxies(nil)

	// Step 2: Initialize the repository
	repo := repository.NewRepository(db)
	competitionRepo := competition.NewCompetitionRepository(repo)
	categoryRepo := category.NewCategoryRepository(repo)
	ropeClimberRepo := ropeclimber.NewRopeClimberRepository(repo)
	timeRepo := timeRepo.NewTimeRepository(repo)
	participationRepo := participation.NewParticipationRepository(repo)
	resultRepo := result.NewResultRepository(repo)

	// Step 3: Initialize the service
	resultService := resultService.NewResultService(resultRepo)
	ropeClimberService := ropeClimberService.NewRopeClimberService(ropeClimberRepo)
	competitionService := competitionService.NewCompetitionService(competitionRepo, categoryRepo, ropeClimberRepo, timeRepo, participationRepo)

	// Step 4: Initialize the handler
	uploadHandler := upload.NewUploadHandler(competitionService)

	// Create a sub-filesystem from the embedded files
	files, err := fs.ReadDir(staticFS, ".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		log.Println(file.Name())
	}
	fsys, err := fs.Sub(staticFS, "static")
	if err != nil {
		log.Fatal(err)
	}
	// Serve the static files
	router.StaticFS("/static", http.FS(fsys))

	router.GET("/logout", func(c *gin.Context) {
		c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
		c.String(http.StatusUnauthorized, "Logging out...")
	})

	router.POST("/upload", BasicAuthMiddleware(), uploadHandler.Upload)
	// pages
	router.GET("/", func(c *gin.Context) {
		seasons := competitionService.GetSeasons()
		renderPartial(c, "competition-list.html", gin.H{
			"Seasons": seasons,
		})
	})

	router.DELETE("/results/:competitionId", BasicAuthMiddleware(), func(c *gin.Context) {
		idStr := c.Param("competitionId")
		id, _ := strconv.ParseInt(idStr, 10, 64)
		competitionService.DeleteCompetition(id)
		seasons := competitionService.GetSeasons()
		renderPartial(c, "management.html", gin.H{
			"Seasons": seasons,
		})
	})

	router.GET("/nomination", func(c *gin.Context) {
		categories := categoryRepo.GetAllCategories()
		year := competitionService.GetSeasons()[0]
		renderPartial(c, "nomination.html", gin.H{
			"Categories": categories,
			"Year":       year,
		})
	})

	router.GET("/nomination-table", func(c *gin.Context) {
		//id := c.Param("id")
		categoryId := c.Query("category")
		category := categoryRepo.FindCategoryById(categoryId)
		timeLimit := os.Getenv(strings.Split(category.CategoryKey, "_")[1] + "_NOMINATION_TIME")
		requiredParticipationCount := os.Getenv(strings.Split(category.CategoryKey, "_")[1] + "_NOMINATION_PARTICIPATION_COUNT")
		requiredParticipationCountInt, _ := strconv.Atoi(requiredParticipationCount)
		timeFloat, _ := strconv.ParseFloat(timeLimit, 64)
		year := competitionService.GetSeasons()[0]
		nominations := resultService.GetNominations(categoryId, strconv.Itoa(int(year)), requiredParticipationCountInt, timeFloat)
		renderPartial(c, "nomination-table.html", gin.H{
			"Nominations": nominations,
		})
	})

	router.GET("/competition-list", func(c *gin.Context) {
		seasons := competitionService.GetSeasons()
		renderPartial(c, "competition-list.html", gin.H{
			"Seasons": seasons,
		})
	})

	router.GET("/rope-climber-competitions/:ropeClimberId", func(c *gin.Context) {
		ropeClimberId := c.Param("ropeClimberId")
		ropeClimberCompetitions := ropeClimberService.GetRopeClimberResults(ropeClimberId)
		renderPartial(c, "rope_climber_competitions.html", gin.H{
			"RopeClimberCompetitions": ropeClimberCompetitions,
		})
	})

	router.GET("/contact", func(c *gin.Context) {
		renderPartial(c, "contact.html", gin.H{})
	})

	router.GET("/results/:competitionId", func(c *gin.Context) {
		competitionId := c.Param("competitionId")
		categories := categoryRepo.GetCategoriesByCompetitionId(competitionId)
		competition := competitionRepo.FindById(competitionId)
		categoryId := c.Query("category")

		var categoryLabel string
		for _, category := range categories {
			if fmt.Sprintf("%d", category.ID) == categoryId { // Convert both to string to match
				categoryLabel = category.Label
				break
			}
		}
		if categoryId == "" {
			categoryId = fmt.Sprint(categories[0].ID)
			categoryLabel = categories[0].Label

		}
		renderPartial(c, "results.html", gin.H{
			"CompetitionID": competitionId,
			"CategoryID":    categoryId,
			"CategoryLabel": categoryLabel,
			"Name":          competition.Name,
			"Categories":    categories,
		})
	})

	router.GET("/top-results", func(c *gin.Context) {
		categories := categoryRepo.GetAllCategories()
		renderPartial(c, "top-results.html", gin.H{
			"Categories": categories,
		})
	})

	router.GET("/management", BasicAuthMiddleware(), func(c *gin.Context) {
		seasons := competitionService.GetSeasons()
		renderPartial(c, "management.html", gin.H{
			"Seasons": seasons,
		})
	})

	// partials
	// season dropdown
	router.GET("/result-table/:competitionId", func(c *gin.Context) {
		competitionId := c.Param("competitionId")
		categoryId := c.Query("category")
		participationResults := resultService.GetResults(competitionId, categoryId)

		renderPartial(c, "result-table.html", gin.H{
			"ParticipationResults": participationResults,
			"TimeCount":            len(participationResults[0].GetTopTimes()),
		})
	})

	router.GET("/top-result-table", func(c *gin.Context) {
		//id := c.Param("id")
		categoryId := c.Query("category")
		participationResults := resultService.GetTopResults(categoryId)
		renderPartial(c, "top-result-table.html", gin.H{
			"TopParticipationResults": participationResults,
		})
	})

	// Dynamic competition list route
	router.GET("/competitions", func(c *gin.Context) {
		seasonParam := c.Query("season")
		seasonInt, err := strconv.Atoi(seasonParam) // Convert to int
		if err != nil {
			c.JSON(400, gin.H{"error": "Invalid season parameter"})
			return
		}
		competitions := competitionService.GetCompetitions(seasonInt)
		showDelete := c.Query("showDelete")
		renderPartial(c, "competitions.html", gin.H{
			"Competitions": competitions,
			"ShowDelete":   showDelete,
		})
	})

	router.Run(":8080")
}

func renderPage(c *gin.Context, templateName string, h gin.H) {
	resultView := c.Query("resultView")
	h["contentTemplate"] = templateName
	h["ResultView"] = resultView
	c.HTML(http.StatusOK, "layout.html", h)
}

func renderPartial(c *gin.Context, templateName string, h gin.H) {
	if c.GetHeader("HX-Request") == "true" {
		c.HTML(http.StatusOK, templateName, h)
		return
	}
	renderPage(c, templateName, h)
}

// Parse templates from the main directory and components subdirectory
func parseTemplates(staticFiles embed.FS) *template.Template {
	tmpl := template.New("") // Create a new template instance

	// Parse main templates
	tmpl = template.Must(template.ParseFS(staticFiles, "templates/*.html"))

	// Parse templates from the components subdirectory
	tmpl = template.Must(tmpl.ParseFS(staticFiles, "templates/components/*.html"))
	return tmpl
}

func BasicAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := os.Getenv("BASIC_AUTH_USERNAME")
		password := os.Getenv("BASIC_AUTH_PASSWORD")
		user, pass, hasAuth := c.Request.BasicAuth()
		if !hasAuth || user != username || pass != password {
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
