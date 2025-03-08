package result

import (
	"github.com/thesimdak/goisos/internal/models"
	"github.com/thesimdak/goisos/internal/repository/result"
)

type ResultService struct {
	ResultRepository *result.ResultRepository
}

func (svc *ResultService) GetResults(competitionId string, categoryId string) []models.ParticipationResult {
	participations := svc.ResultRepository.FindResultsBycompetitionIdAndCategoryId(competitionId, categoryId)
	return participations
}

func NewResultService(resultRepository *result.ResultRepository,
) *ResultService {
	return &ResultService{ResultRepository: resultRepository}
}
