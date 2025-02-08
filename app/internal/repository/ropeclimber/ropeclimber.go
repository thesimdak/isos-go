package ropeclimber

import (
	"log"

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
