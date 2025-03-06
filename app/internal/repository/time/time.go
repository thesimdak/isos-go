package time

import (
	"log"

	"github.com/thesimdak/goisos/internal/models"
	"github.com/thesimdak/goisos/internal/repository"
)

type TimeRepository struct {
	*repository.Repository
}

func NewTimeRepository(repo *repository.Repository) *TimeRepository {
	return &TimeRepository{Repository: repo}
}

func (repo *TimeRepository) DeleteByCompetitionIdAndCategoryIdAndRopeClimberId(competitionId int64, categoryId int64, ropeClimberId int64) {
	query := `DELETE t FROM time t JOIN participation p on t.participation_id = p.id WHERE p.competition_id = ? and p.category_id = ? and p.rope_climber_id = ?`
	_, err := repo.DB.Exec(
		query,
		competitionId,
		categoryId,
		ropeClimberId,
	)
	if err != nil {
		log.Printf("Error inserting time: %v", err)
		panic("Error inserting time")
	}
}

func (repo *TimeRepository) DeleteByCompetitionId(competitionId int64) {
	query := `DELETE t FROM time t JOIN participation p on t.participation_id = p.id WHERE p.competition_id = ?`
	_, err := repo.DB.Exec(
		query,
		competitionId,
	)
	if err != nil {
		log.Printf("Error inserting time: %v", err)
		panic("Error inserting time")
	}
}

func (repo *TimeRepository) SaveTimes(participationId int64, times []*models.Time) {
	for _, t := range times {
		query := `INSERT INTO time (
			participation_id, time, round
		) VALUES (?, ?, ?)`

		_, err := repo.DB.Exec(
			query,
			participationId,
			t.Time,
			t.Round,
		)
		if err != nil {
			log.Printf("Error inserting time: %v", err)
			panic("Error inserting time")
		}
	}
}
