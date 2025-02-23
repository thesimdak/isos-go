package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/thesimdak/goisos/internal/handlers/upload"
	models "github.com/thesimdak/goisos/internal/models"
	"github.com/thesimdak/goisos/internal/repository"
	"github.com/thesimdak/goisos/internal/repository/category"
	"github.com/thesimdak/goisos/internal/repository/competition"
	"github.com/thesimdak/goisos/internal/repository/participation"
	"github.com/thesimdak/goisos/internal/repository/ropeclimber"
	timeRepo "github.com/thesimdak/goisos/internal/repository/time"
	competitionService "github.com/thesimdak/goisos/internal/services/competition"
)

func Initialize(db *sql.DB) {
	router := gin.Default()

	// Step 2: Initialize the repository
	repo := repository.NewRepository(db)
	competitionRepo := competition.NewCompetitionRepository(repo)
	categoryRepo := category.NewCategoryRepository(repo)
	ropeClimberRepo := ropeclimber.NewRopeClimberRepository(repo)
	timeRepo := timeRepo.NewTimeRepository(repo)
	participationRepo := participation.NewParticipationRepository(repo)

	// Step 3: Initialize the service
	competitionService := competitionService.NewCompetitionService(competitionRepo, categoryRepo, ropeClimberRepo, timeRepo, participationRepo)

	// Step 4: Initialize the handler
	uploadHandler := upload.NewUploadHandler(competitionService)
	router.Static("/static", "./static")

	router.POST("/upload", uploadHandler.Upload)
	// pages
	router.GET("/", func(c *gin.Context) {
		seasons := competitionService.GetSeasons()
		renderPartial(c, "competition-list.html", gin.H{
			"Seasons": seasons,
		})
	})
	router.GET("/competition-list", func(c *gin.Context) {
		//seasons := competitionService.GetSeasons()
		seasons := []int16{2010, 2011, 2012, 2013, 2014, 2015, 2016, 2017, 2018, 2019, 2020, 2021, 2022, 2023, 2024, 2025}
		renderPartial(c, "competition-list.html", gin.H{
			"Seasons": seasons,
		})
	})
	router.GET("/results/:competitionId", func(c *gin.Context) {
		//id := c.Param("id")
		var categories []models.Category
		categories = append(categories, models.Category{ID: 122, Label: "Muzi"})
		categories = append(categories, models.Category{ID: 123, Label: "Zeny"})
		categories = append(categories, models.Category{ID: 124, Label: "Dorostenci"})
		competitionId, _ := c.Params.Get("competitionId")
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
			"Name":          "Memorial Bedricha Supcika",
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

	router.GET("/management", func(c *gin.Context) {
		seasons := []int16{2010, 2011, 2012, 2013, 2014, 2015, 2016, 2017, 2018, 2019, 2020, 2021, 2022, 2023, 2024, 2025}
		renderPartial(c, "management.html", gin.H{
			"Seasons": seasons,
		})
	})

	// partials
	// season dropdown
	router.GET("/result-table/:competitionId", func(c *gin.Context) {
		//id := c.Param("id")
		var participationResults []models.ParticipationResult
		participationResults = append(participationResults, models.ParticipationResult{Rank: 1, Name: "Jiri Novak", YearOfBirth: "1992", Organization: "Sokol Liben", Time1: "9.56", Time2: "9.56", Time3: "-", Time4: "5.56", Top: "5.56"})
		participationResults = append(participationResults, models.ParticipationResult{Rank: 2, Name: "Martin Simon", YearOfBirth: "1990", Organization: "Sokol Liben", Time1: "11.56", Time2: "9.56", Time3: "-", Time4: "5.56", Top: "11.56"})
		renderPartial(c, "result-table.html", gin.H{
			"ParticipationResults": participationResults,
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
		var competitions []models.Competition
		competitions = append(competitions, models.Competition{ID: 122, Name: "Memorial Bedricha Supcika 2024", Date: time.Now()})
		competitions = append(competitions, models.Competition{ID: 123, Name: "Modransky Tarzan", Date: time.Now()})
		showDelete := c.Query("showDelete")
		c.HTML(http.StatusOK, "competitions.html", gin.H{
			"Competitions": competitions,
			"ShowDelete":   showDelete,
		})
	})

	// Load templates
	router.SetHTMLTemplate(parseTemplates())
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
