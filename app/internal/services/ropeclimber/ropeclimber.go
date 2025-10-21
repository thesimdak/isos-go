package ropeclimber

import (
	"fmt"

	"github.com/thesimdak/goisos/internal/models"
	"github.com/thesimdak/goisos/internal/repository/ropeclimber"
)

type RopeClimberService struct {
	RopeClimberRepository *ropeclimber.RopeClimberRepository
}

func (svc *RopeClimberService) GetRopeClimberResults(roleClimberId string) []models.RopeClimberCompetition {
	ropeClimberCompetitions := svc.RopeClimberRepository.FindRopeClimberCompetition(roleClimberId)

	setTopTimeAndReplaceInvalid(ropeClimberCompetitions)
	return ropeClimberCompetitions
}

func NewRopeClimberService(ropeClimberRepository *ropeclimber.RopeClimberRepository) *RopeClimberService {
	return &RopeClimberService{RopeClimberRepository: ropeClimberRepository}
}

func setTopTimeAndReplaceInvalid(ropeClimberCompetition []models.RopeClimberCompetition) {
	for i := 0; i < len(ropeClimberCompetition); i++ {
		formattedNum := fmt.Sprintf("%.2f", ropeClimberCompetition[i].GetTopTime())
		if formattedNum == "999.00" {
			dash := "-"
			ropeClimberCompetition[i].Top = &dash
		} else {
			ropeClimberCompetition[i].Top = &formattedNum
		}

		replaceInvalidClimbeNumber(ropeClimberCompetition[i].Time1)
		replaceInvalidClimbeNumber(ropeClimberCompetition[i].Time2)
		replaceInvalidClimbeNumber(ropeClimberCompetition[i].Time3)
		replaceInvalidClimbeNumber(ropeClimberCompetition[i].Time4)
	}
}

func replaceInvalidClimbeNumber(time *string) {
	if time != nil && *time == "999.00" {
		*time = "-"
	}
}
