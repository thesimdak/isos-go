package participation

import (
	"log"

	"github.com/thesimdak/goisos/internal/models"
	"github.com/thesimdak/goisos/internal/repository"
)

type ParticipationRepository struct {
	*repository.Repository
}

func NewParticipationRepository(repo *repository.Repository) *ParticipationRepository {
	return &ParticipationRepository{Repository: repo}
}

func (repo *ParticipationRepository) DeleteByCompetitionId(id int64) {
	query := `DELETE FROM participation WHERE competition_id = ?`

	repo.DB.Exec(
		query,
		id,
	)
}

func (repo *ParticipationRepository) DeleteByCompetitionIdAndCategoryIdAndRopeClimberId(p *models.Participation) {
	query := `DELETE FROM participation WHERE competition_id = ? and category_id = ? and rope_climber_id = ?`

	repo.DB.Exec(
		query,
		p.Competition.ID,
		p.Category.ID,
		p.RopeClimber.ID,
	)
}

func (repo *ParticipationRepository) InsertParticipation(p *models.Participation) *models.Participation {
	query := `INSERT INTO participation (
			category_id, competition_id, organization, rope_climber_id
		) VALUES (?, ?, ?, ?)`

	result, err := repo.DB.Exec(
		query,
		p.Category.ID,
		p.Competition.ID,
		p.Organization,
		p.RopeClimber.ID,
	)
	if err != nil {
		log.Printf("Error inserting participation: %v", err)
		panic("Error inserting participation")
	}
	id, _ := result.LastInsertId()
	p.ID = id
	return p
}
