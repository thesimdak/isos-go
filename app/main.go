package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/thesimdak/goisos/internal/db"
	"github.com/thesimdak/goisos/internal/handlers"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Warning: .env file not found, using default values")
	}
	db := db.InitializeDB() // Initialize your database
	defer db.Close()

	handlers.Initialize(db)

}
