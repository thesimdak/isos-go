package main

import (
	"embed"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/thesimdak/goisos/internal/db"
	"github.com/thesimdak/goisos/internal/handlers"
)

//go:embed static/*
//go:embed templates/*.html
//go:embed templates/components/*.html
var staticFiles embed.FS

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found, using default values")
	}

	db := db.InitializeDB() // Initialize your database
	defer db.Close()

	handlers.Initialize(db, staticFiles)

}
