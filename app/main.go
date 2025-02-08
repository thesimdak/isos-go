package main

import (
	"github.com/thesimdak/goisos/internal/db"
	"github.com/thesimdak/goisos/internal/handlers"
)

func main() {
	conn := db.InitializeDB()
	db := db.InitializeDB() // Initialize your database
	defer db.Close()

	handlers.Initialize(conn)

}
