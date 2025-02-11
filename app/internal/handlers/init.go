package handlers

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thesimdak/goisos/internal/handlers/upload"
	"github.com/thesimdak/goisos/internal/repository"
	"github.com/thesimdak/goisos/internal/repository/category"
	"github.com/thesimdak/goisos/internal/repository/competition"
	"github.com/thesimdak/goisos/internal/repository/participation"
	"github.com/thesimdak/goisos/internal/repository/ropeclimber"
	"github.com/thesimdak/goisos/internal/repository/time"
	competitionService "github.com/thesimdak/goisos/internal/services/competition"
)

func Initialize(db *sql.DB) {
	router := gin.Default()

	// Step 2: Initialize the repository
	repo := repository.NewRepository(db)
	competitionRepo := competition.NewCompetitionRepository(repo)
	categoryRepo := category.NewCategoryRepository(repo)
	ropeClimberRepo := ropeclimber.NewRopeClimberRepository(repo)
	timeRepo := time.NewTimeRepository(repo)
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
	router.GET("/top-results", func(c *gin.Context) {
		renderPartial(c, "top-results.html", gin.H{})
	})
	router.GET("/management", func(c *gin.Context) {
		renderPartial(c, "management.html", gin.H{})
	})

	// partials
	// season dropdown
	router.GET("/seasons", func(c *gin.Context) {
		renderPartial(c, "seasons.html", nil)
	})

	// Load templates
	router.SetHTMLTemplate(parseTemplates())

	// Dynamic competition list route
	router.GET("/competitions", func(c *gin.Context) {
		var competitions []string
		competitions = append(competitions, "Memorial Bedricha Supcika")
		competitions = append(competitions, "Mistrovstvi Ceske Republiky")
		c.HTML(http.StatusOK, "competitions.html", gin.H{
			"Competitions": competitions,
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
