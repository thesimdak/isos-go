package ropeclimber

import (
	"log"
	"time"

	"github.com/thesimdak/goisos/internal/models"
	"github.com/thesimdak/goisos/internal/repository"
)

type RopeClimberRepository struct {
	*repository.Repository
}

func NewRopeClimberRepository(repo *repository.Repository) *RopeClimberRepository {
	return &RopeClimberRepository{Repository: repo}
}

func (repo *RopeClimberRepository) EnsureRopeClimber(rc *models.RopeClimber) *models.RopeClimber {
	ropeClimber := repo.FindRopeClimber(rc)
	if ropeClimber == nil {
		return repo.InsertRopeClimber(rc)
	}

	return ropeClimber
}

func (repo *RopeClimberRepository) InsertRopeClimber(rc *models.RopeClimber) *models.RopeClimber {
	query := `INSERT INTO rope_climber (
			first_name, last_name, year_of_birth
		) VALUES (?, ?, ?)`

	result, err := repo.DB.Exec(
		query,
		rc.FirstName,
		rc.LastName,
		rc.YearOfBirth,
	)
	if err != nil {
		log.Printf("Error inserting rope climber: %v", err)
		panic("Error inserting rope climber")
	}
	id, _ := result.LastInsertId()
	rc.ID = id
	return rc
}

func (repo *RopeClimberRepository) FindRopeClimber(rc *models.RopeClimber) *models.RopeClimber {
	query := `SELECT id, first_name, last_name, year_of_birth FROM rope_climber r WHERE r.first_name = ? and r.last_name = ? and r.year_of_birth = ?`

	row := repo.DB.QueryRow(query, rc.FirstName, rc.LastName, rc.YearOfBirth)

	var ropeClimber models.RopeClimber

	err := row.Scan(
		&ropeClimber.ID,
		&ropeClimber.FirstName,
		&ropeClimber.LastName,
		&ropeClimber.YearOfBirth,
	)
	if err != nil {
		return nil
	}
	return &ropeClimber
}

func (repo *RopeClimberRepository) FindRopeClimberCompetition(ropeclimberId string) []models.RopeClimberCompetition {
	query := `SELECT rc.id, 
					 CONCAT(rc.last_name, ', ', rc.first_name) AS name,
					 rc.year_of_birth, 
					 p.organization, 
					 cat.label,
					 c.name, 
					 c.date, 
					 FORMAT(t1.time, 2), 
					 FORMAT(t2.time, 2), 
					 FORMAT(t3.time, 2), 
					 FORMAT(t4.time, 2) FROM rope_climber rc JOIN participation p ON p.rope_climber_id = rc.id 
					 										JOIN category cat ON cat.id = p.category_id 	
LEFT JOIN time t1 ON t1.participation_id = p.id AND t1.round = 1
    LEFT JOIN time t2 ON t2.participation_id = p.id AND t2.round = 2
    LEFT JOIN time t3 ON t3.participation_id = p.id AND t3.round = 3
    LEFT JOIN time t4 ON t4.participation_id = p.id AND t4.round = 4
JOIN competition c ON c.id = p.competition_id WHERE rc.id = ? ORDER BY c.date desc`

	rows, err := repo.DB.Query(query, ropeclimberId)

	if err != nil {
		return make([]models.RopeClimberCompetition, 0)
	}
	defer rows.Close() // Ensure the rows are closed after we are done

	var results []models.RopeClimberCompetition

	// Iterate over the rows to populate the slice of competitions
	for rows.Next() {
		var result models.RopeClimberCompetition
		var dateBytes []byte
		err := rows.Scan(
			&result.ID,
			&result.Name,
			&result.YearOfBirth,
			&result.Organization,
			&result.Category,
			&result.CompetitionName,
			&dateBytes,
			&result.Time1,
			&result.Time2,
			&result.Time3,
			&result.Time4,
		)
		if err != nil {
			return make([]models.RopeClimberCompetition, 0) // Return empty
		}
		parsedDate, err := time.Parse("2006-01-02", string(dateBytes))

		result.Date = &parsedDate
		// Add the competition to the slice
		results = append(results, result)
	}

	return results
}
