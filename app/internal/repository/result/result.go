package result

import (
	"github.com/thesimdak/goisos/internal/models"
	"github.com/thesimdak/goisos/internal/repository"
)

type ResultRepository struct {
	*repository.Repository
}

func (repo *ResultRepository) FindNominationsByCategoryId(categoryId string, year string) []models.Nomination {
	query := `WITH MinTimes AS (
    SELECT
        p.id AS participation_id,
        MIN(t.time) AS min_time
    FROM
        participation p
        JOIN time t ON t.participation_id = p.id
        JOIN competition c ON c.id = p.competition_id
    WHERE p.category_id = 3 AND YEAR(c.date) = 2025
    GROUP BY
        p.id
),
RankedParticipants AS (
    SELECT
        rc.id as id,
        CONCAT(rc.last_name, ', ', rc.first_name) AS name,
        rc.year_of_birth,
        p.organization,
        c.name AS c_name,
        mt.min_time AS top,
        ROW_NUMBER() OVER (
            PARTITION BY rc.last_name, rc.first_name, rc.year_of_birth
            ORDER BY mt.min_time, c.date
        ) AS rn
    FROM
        participation p
        JOIN rope_climber rc ON rc.id = p.rope_climber_id
        JOIN MinTimes mt ON mt.participation_id = p.id
        JOIN competition c ON c.id = p.competition_id
    GROUP BY
        rc.id, rc.last_name, rc.first_name, rc.year_of_birth, p.organization, mt.min_time, c.date, c.name
),
ParticipationCount AS (
    SELECT
        rc.id as rc_id, COUNT(*) as count
    FROM
        participation p
        JOIN rope_climber rc ON rc.id = p.rope_climber_id
        JOIN competition c ON c.id = p.competition_id AND YEAR(c.date) = 2025
    GROUP BY
        rc.id
)
SELECT
    rp.name,
    rp.year_of_birth,
    rp.organization,
    rp.c_name,
    pc.count,
    FORMAT(rp.top, 2) AS formatted_top
FROM
    RankedParticipants rp
    JOIN ParticipationCount pc ON pc.rc_id = rp.id
WHERE
    rp.rn = 1
ORDER BY
    rp.top;

			`

	rows, err := repo.DB.Query(query, categoryId)
	if err != nil {
		return make([]models.Nomination, 0) // Return an error if the query fails
	}
	defer rows.Close() // Ensure the rows are closed after we are done

	var results []models.Nomination

	// Iterate over the rows to populate the slice of competitions
	for rows.Next() {
		var result models.Nomination
		err := rows.Scan(
			&result.Name,
			&result.YearOfBirth,
			&result.Organization,
			&result.CompetitionName,
			&result.ParticipationCount,
			&result.BestTime,
		)
		if err != nil {
			return make([]models.Nomination, 0) // Return empty
		}

		// Add the competition to the slice
		results = append(results, result)
	}

	return results
}

func (repo *ResultRepository) FindTopResultsByCategoryId(categoryId string) []models.TopParticipationResults {
	query := `WITH MinTimes AS (
    SELECT
        p.id AS participation_id,
        MIN(t.time) AS min_time
    FROM
        participation p
        JOIN time t ON t.participation_id = p.id
    WHERE p.category_id = ?
    GROUP BY
        p.id
	),
	RankedParticipants AS (
		SELECT
			CONCAT(rc.last_name, ', ', rc.first_name) AS name,
			rc.year_of_birth,
			p.organization,
			c.name AS c_name,
			mt.min_time AS top,
			ROW_NUMBER() OVER (
				PARTITION BY rc.last_name, rc.first_name, rc.year_of_birth
				ORDER BY mt.min_time, c.date
			) AS rn
		FROM
			participation p
			JOIN rope_climber rc ON rc.id = p.rope_climber_id
			JOIN MinTimes mt ON mt.participation_id = p.id
			JOIN competition c ON c.id = p.competition_id
		GROUP BY
			rc.last_name, rc.first_name, rc.year_of_birth, p.organization, mt.min_time, c.date, c.name
		)
		SELECT
			name,
			year_of_birth,
			organization,
			c_name,
			FORMAT(top, 2)
		FROM
			RankedParticipants
		WHERE
			rn = 1
		ORDER BY
			top
		`

	rows, err := repo.DB.Query(query, categoryId)
	if err != nil {
		return make([]models.TopParticipationResults, 0) // Return an error if the query fails
	}
	defer rows.Close() // Ensure the rows are closed after we are done

	var results []models.TopParticipationResults

	// Iterate over the rows to populate the slice of competitions
	for rows.Next() {
		var result models.TopParticipationResults
		err := rows.Scan(
			&result.Name,
			&result.YearOfBirth,
			&result.Organization,
			&result.CompetitionName,
			&result.Top,
		)
		if err != nil {
			return make([]models.TopParticipationResults, 0) // Return empty
		}

		// Add the competition to the slice
		results = append(results, result)
	}

	return results
}

func NewResultRepository(repo *repository.Repository) *ResultRepository {
	return &ResultRepository{Repository: repo}
}

func (repo *ResultRepository) FindResultsBycompetitionIdAndCategoryId(competitionId string, categoryId string) []models.ParticipationResult {
	query := `SELECT CONCAT(rc.last_name, ', ', rc.first_name) AS name, rc.year_of_birth, p.organization, FORMAT(t1.time, 2), FORMAT(t2.time, 2), FORMAT(t3.time, 2), FORMAT(t4.time, 2) FROM participation p 
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
