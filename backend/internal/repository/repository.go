package repository

import "database/sql"

// Repository struct that holds the DB connection
type Repository struct {
	DB *sql.DB
}

// NewRepository creates a new Repository instance
func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}
