package category

import (
	"log"

	"github.com/thesimdak/goisos/internal/models"
	"github.com/thesimdak/goisos/internal/repository"
)

type CategoryRepository struct {
	*repository.Repository
}

func (repo *CategoryRepository) GetAllCategories() any {
	query := `
        SELECT distinct cat.id, cat.label
        	FROM category cat
				JOIN participation p on p.category_id = cat.id`

	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil // Return an error if the query fails
	}
	defer rows.Close() // Ensure the rows are closed after we are done

	var categories []models.Category

	// Iterate over the rows to populate the slice of competitions
	for rows.Next() {
		var category models.Category

		err := rows.Scan(
			&category.ID,
			&category.Label,
		)
		if err != nil {
			return make([]models.Category, 0)
		}

		// Add the competition to the slice
		categories = append(categories, category)
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		return make([]models.Category, 0)
	}

	return categories
}

func (repo *CategoryRepository) GetCategoriesByCompetitionId(competitionId string) []models.Category {
	query := `
        SELECT distinct cat.id, cat.label
        	FROM category cat
				JOIN participation p on p.category_id = cat.id where p.competition_id = ?`

	rows, err := repo.DB.Query(query, competitionId)
	if err != nil {
		return nil // Return an error if the query fails
	}
	defer rows.Close() // Ensure the rows are closed after we are done

	var categories []models.Category

	// Iterate over the rows to populate the slice of competitions
	for rows.Next() {
		var category models.Category

		err := rows.Scan(
			&category.ID,
			&category.Label,
		)
		if err != nil {
			return make([]models.Category, 0)
		}

		// Add the competition to the slice
		categories = append(categories, category)
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		return make([]models.Category, 0)
	}

	return categories
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

func (repo *CategoryRepository) FindCategoryById(categoryId string) *models.Category {
	query := `SELECT id, category_key, label, rope_length  FROM category WHERE category.id = ?`

	row := repo.DB.QueryRow(query, categoryId)

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
