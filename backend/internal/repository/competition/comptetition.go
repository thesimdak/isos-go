package competition

import (
	"log"
	"time"

	"github.com/thesimdak/goisos/internal/models"
	"github.com/thesimdak/goisos/internal/repository"
)

type CompetitionRepository struct {
	*repository.Repository
}

// NewCompetitionRepository creates a new CompetitionRepository instance
func NewCompetitionRepository(repo *repository.Repository) *CompetitionRepository {
	return &CompetitionRepository{Repository: repo}
}

func (repo *CompetitionRepository) SaveCompetition(competition *models.Competition) *models.Competition {
	query := `
		INSERT INTO competition (
			competition_name, name, date, place, jugde, sensor_installation, starter, type
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := repo.DB.Exec(
		query,
		competition.CompetitionName,
		competition.Name,
		competition.Date,
		competition.Place,
		competition.Judge,
		competition.SensorInstallation,
		competition.Starter,
		competition.Type,
	)
	if err != nil {
		log.Printf("Error inserting competition: %v", err)
	}
	id, _ := result.LastInsertId()
	competition.ID = id
	return competition
}

func (repo *CompetitionRepository) FindCompetition(name string, date time.Time) *models.Competition {
	query := `
        SELECT id, competition_name, name, date, place, jugde, sensor_installation, starter, type
        FROM competition
        WHERE name = ? AND date = ?`

	row := repo.DB.QueryRow(query, name, date)

	var competition models.Competition

	var dateBytes []byte
	err := row.Scan(
		&competition.ID,
		&competition.CompetitionName,
		&competition.Name,
		&dateBytes,
		&competition.Place,
		&competition.Judge,
		&competition.SensorInstallation,
		&competition.Starter,
		&competition.Type,
	)
	if err != nil {
		return nil
	}

	// Parse the date manually if it's in []byte format
	parsedDate, err := time.Parse("2006-01-02", string(dateBytes))

	competition.Date = parsedDate // Assign the parsed date to the competition struct

	// Handle any errors from scanning
	if err != nil {
		return nil
	}

	return &competition
}
