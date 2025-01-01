package category

import (
	"log"

	"github.com/thesimdak/goisos/internal/models"
	"github.com/thesimdak/goisos/internal/repository"
)

type CategoryRepository struct {
	*repository.Repository
}

// NewCompetitionRepository creates a new CompetitionRepository instance
func NewCategoryRepository(repo *repository.Repository) *CategoryRepository {
	return &CategoryRepository{Repository: repo}
}

func (repo *CategoryRepository) EnsureCategory(categoryKey string, label string, ropeLength float32) *models.Category {
	category := repo.FindCategory(categoryKey)
	if category == nil {
		return repo.InsertCategory(categoryKey, label, ropeLength)
	}

	return category
}

func (repo *CategoryRepository) InsertCategory(categoryKey string, label string, ropeLength float32) *models.Category {
	query := `INSERT INTO category (
			category_key, label, rope_length
		) VALUES (?, ?, ?)`

	result, err := repo.DB.Exec(
		query,
		categoryKey,
		label,
		ropeLength,
	)
	if err != nil {
		log.Printf("Error inserting category: %v", err)
		panic("Error inserting category")
	}
	id, _ := result.LastInsertId()
	return &models.Category{ID: id, CategoryKey: categoryKey, Label: label, RopeLength: ropeLength}
}

func (repo *CategoryRepository) FindCategory(categoryKey string) *models.Category {
	query := `SELECT id, category_key, label, rope_length  FROM category WHERE category.category_key = ?`

	row := repo.DB.QueryRow(query, categoryKey)

	var category models.Category

	err := row.Scan(
		&category.ID,
		&category.CategoryKey,
		&category.Label,
		&category.RopeLength,
	)
	if err != nil {
		return nil
	}
	return &category
}
