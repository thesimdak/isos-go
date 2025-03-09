package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/thesimdak/goisos/internal/handlers/upload"
	models "github.com/thesimdak/goisos/internal/models"
	"github.com/thesimdak/goisos/internal/repository"
	"github.com/thesimdak/goisos/internal/repository/category"
	"github.com/thesimdak/goisos/internal/repository/competition"
	"github.com/thesimdak/goisos/internal/repository/participation"
	"github.com/thesimdak/goisos/internal/repository/result"
	"github.com/thesimdak/goisos/internal/repository/ropeclimber"
	timeRepo "github.com/thesimdak/goisos/internal/repository/time"
	competitionService "github.com/thesimdak/goisos/internal/services/competition"
	resultService "github.com/thesimdak/goisos/internal/services/result"
)

func Initialize(db *sql.DB) {
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
	router.SetHTMLTemplate(parseTemplates())
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
	competitionService := competitionService.NewCompetitionService(competitionRepo, categoryRepo, ropeClimberRepo, timeRepo, participationRepo)

	// Step 4: Initialize the handler
	uploadHandler := upload.NewUploadHandler(competitionService)
	router.Static("/static", "./static")
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

	router.GET("/competition-list", func(c *gin.Context) {
		seasons := competitionService.GetSeasons()
		renderPartial(c, "competition-list.html", gin.H{
			"Seasons": seasons,
		})
	})

	router.GET("/results/:competitionId", func(c *gin.Context) {
		//id := c.Param("id")
		competitionId := c.Param("competitionId")
		categories := categoryRepo.GetCategoriesByCompetitionId(competitionId)
		competition := competitionRepo.FindById(competitionId)
		categoryId := c.Query("categoryId")

		var categoryLabel string
		for _, category := range categories {
			if fmt.Sprintf("%d", category.ID) == categoryId { // Convert both to string to match
				categoryLabel = category.Label
				break
			}
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
		var categories []models.Category
		categories = append(categories, models.Category{ID: 122, Label: "Muzi"})
		categories = append(categories, models.Category{ID: 123, Label: "Zeny"})
		categories = append(categories, models.Category{ID: 124, Label: "Dorostenci"})
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
		categoryId := c.Query("categoryId")
		participationResults := resultService.GetResults(competitionId, categoryId)

		renderPartial(c, "result-table.html", gin.H{
			"ParticipationResults": participationResults,
			"TimeCount":            len(participationResults[0].GetTopTimes()),
		})
	})

	router.GET("/top-result-table", func(c *gin.Context) {
		//id := c.Param("id")
		var topParticipationResults []models.TopParticipationResults
		topParticipationResults = append(topParticipationResults, models.TopParticipationResults{Rank: 1, Name: "Jiri Novak", YearOfBirth: "1992", Organization: "Sokol Liben", CompetitionName: "Memorial Bedricha Supcika", Top: "5.56"})
		topParticipationResults = append(topParticipationResults, models.TopParticipationResults{Rank: 2, Name: "Martin Simon", YearOfBirth: "1990", Organization: "Sokol Liben", CompetitionName: "Pisecky Splhavec", Top: "5.56"})
		renderPartial(c, "top-result-table.html", gin.H{
			"TopParticipationResults": topParticipationResults,
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
	h["contentTemplate"] = templateName
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
func parseTemplates() *template.Template {
	tmpl := template.New("") // Create a new template instance

	// Parse templates from the main directory
	tmpl = template.Must(tmpl.ParseGlob("templates/*.html"))

	// Parse templates from the components subdirectory
	tmpl = template.Must(tmpl.ParseGlob("templates/components/*.html"))

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
