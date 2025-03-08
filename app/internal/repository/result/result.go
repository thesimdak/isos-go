package result

import (
	"github.com/thesimdak/goisos/internal/models"
	"github.com/thesimdak/goisos/internal/repository"
)

type ResultRepository struct {
	*repository.Repository
}

func NewResultRepository(repo *repository.Repository) *ResultRepository {
	return &ResultRepository{Repository: repo}
}

func (repo *ResultRepository) FindResultsBycompetitionIdAndCategoryId(competitionId string, categoryId string) []models.ParticipationResult {
	query := `SELECT CONCAT(rc.last_name, ', ', rc.first_name) AS name, rc.year_of_birth, p.organization, t1.time, t2.time, t3.time, t4.time FROM participation p 
	JOIN rope_climber rc ON rc.id = p.rope_climber_id
	LEFT JOIN time t1 ON t1.participation_id = p.id and t1.round = 1
    LEFT JOIN time t2 ON t2.participation_id = p.id and t2.round = 2
    LEFT JOIN time t3 ON t3.participation_id = p.id and t3.round = 3
    LEFT JOIN time t4 ON t4.participation_id = p.id and t4.round = 4 where p.competition_id = ? and p.category_id = ?`

	rows, err := repo.DB.Query(query, competitionId, categoryId)
	if err != nil {
		return make([]models.ParticipationResult, 0) // Return an error if the query fails
	}
	defer rows.Close() // Ensure the rows are closed after we are done

	var results []models.ParticipationResult

	// Iterate over the rows to populate the slice of competitions
	for rows.Next() {
		var result models.ParticipationResult
		err := rows.Scan(
			&result.Name,
			&result.YearOfBirth,
			&result.Organization,
			&result.Time1,
			&result.Time2,
			&result.Time3,
			&result.Time4,
		)
		if err != nil {
			return make([]models.ParticipationResult, 0) // Return empty
		}

		// Add the competition to the slice
		results = append(results, result)
	}

	return results
}
