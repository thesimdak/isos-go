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

func (repo *CompetitionRepository) GetSeasons() ([]int16, error) {
	query := `SELECT DISTINCT YEAR(date) FROM competition ORDER BY YEAR(date) DESC`

	// Execute the query
	rows, err := repo.DB.Query(query)
	if err != nil {
		return nil, err // Handle query execution error
	}
	defer rows.Close()

	// Slice to store the seasons (years)
	var seasons []int16

	// Iterate through the rows and scan the years
	for rows.Next() {
		var year int16
		if err := rows.Scan(&year); err != nil {
			return nil, err // Handle scanning error
		}
		seasons = append(seasons, year)
	}

	// Check for errors during row iteration
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return seasons, nil
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

func (repo *CompetitionRepository) FindAllCompetitionsBySeason(season int) ([]models.Competition, error) {
	query := `
        SELECT id, competition_name, name, date, place, jugde, sensor_installation, starter, type
        FROM competition where YEAR(date) = ?`

	rows, err := repo.DB.Query(query, season)
	if err != nil {
		return nil, err // Return an error if the query fails
	}
	defer rows.Close() // Ensure the rows are closed after we are done

	var competitions []models.Competition

	// Iterate over the rows to populate the slice of competitions
	for rows.Next() {
		var competition models.Competition
		var dateBytes []byte

		err := rows.Scan(
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
			return nil, err // Return an error if scanning fails
		}

		// Parse the date manually if it's in []byte format
		parsedDate, err := time.Parse("2006-01-02", string(dateBytes))
		if err != nil {
			return nil, err // Return an error if parsing the date fails
		}
		competition.Date = parsedDate // Assign the parsed date to the competition struct

		// Add the competition to the slice
		competitions = append(competitions, competition)
	}

	// Check for any errors encountered during iteration
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return competitions, nil
}
