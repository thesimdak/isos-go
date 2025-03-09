package result

import (
	"fmt"
	"sort"

	"github.com/thesimdak/goisos/internal/models"
	"github.com/thesimdak/goisos/internal/repository/result"
)

type ResultService struct {
	ResultRepository *result.ResultRepository
}

func (svc *ResultService) GetResults(competitionId string, categoryId string) []models.ParticipationResult {
	participations := svc.ResultRepository.FindResultsBycompetitionIdAndCategoryId(competitionId, categoryId)
	sortParticipationResults(participations)
	setRankAndTopTime(participations)
	return participations
}

func NewResultService(resultRepository *result.ResultRepository,
) *ResultService {
	return &ResultService{ResultRepository: resultRepository}
}

func sortParticipationResults(results []models.ParticipationResult) {
	sort.Sort(models.ByTopTime(results))
}

func setRankAndTopTime(results []models.ParticipationResult) {
	for i := 0; i < len(results); i++ {
		formattedNum := fmt.Sprintf("%.2f", results[i].GetTopTimes()[0])
		results[i].Top = &formattedNum
		compareAndSetRank(results, i)
		replaceInvalidClimbeNumber(results[i].Time1)
		replaceInvalidClimbeNumber(results[i].Time2)
		replaceInvalidClimbeNumber(results[i].Time3)
		replaceInvalidClimbeNumber(results[i].Time4)
	}
}

func compareAndSetRank(results []models.ParticipationResult, i int) {
	if i == 0 {
		results[i].Rank = i + 1
		return
	}
	previousTimes := results[i-1].GetTopTimes()
	times := results[i].GetTopTimes()
	keepPreviousRank := true
	for t := 0; t < len(previousTimes); t++ {
		if times[t] != previousTimes[t] {
			keepPreviousRank = false
			break
		}
	}
	if keepPreviousRank {
		results[i].Rank = results[i-1].Rank
		return
	}
	results[i].Rank = i + 1
}

func replaceInvalidClimbeNumber(time *string) {
	if time != nil && *time == "999" {
		*time = "-"
	}
}
