package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // Replace with your database driver
)

// InitializeDB initializes and returns a database connection.
func InitializeDB() *sql.DB {
	dbConnection := os.Getenv("DB_CONNECTION")
	dsn := dbConnection // Update with your DB details
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Ensure the connection is available
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	log.Println("Database connection established successfully.")
	return db
}
