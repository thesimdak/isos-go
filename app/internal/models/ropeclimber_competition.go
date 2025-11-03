package models

import (
	"sort"
	"strconv"
	"time"
)

type RopeClimberCompetition struct {
	ID              *int64
	Name            *string
	YearOfBirth     *int16
	Date            *time.Time
	Organization    *string
	Category        *string
	CompetitionName *string
	Time1           *string
	Time2           *string
	Time3           *string
	Time4           *string
	Top             *string
}

func (c RopeClimberCompetition) FormattedDate() string {
	return c.Date.Format("02.01.2006") // Format as DD.MM.YYYY
}

func (r RopeClimberCompetition) GetTopTime() float64 {
	times := make([]float64, 0, 4)

	if r.Time1 != nil {
		if time1, err := strconv.ParseFloat(*r.Time1, 64); err == nil {
			times = append(times, time1)
		}
	}
	if r.Time2 != nil {
		if time2, err := strconv.ParseFloat(*r.Time2, 64); err == nil {
			times = append(times, time2)
		}
	}
	if r.Time3 != nil {
		if time3, err := strconv.ParseFloat(*r.Time3, 64); err == nil {
			times = append(times, time3)
		}
	}
	if r.Time4 != nil {
		if time4, err := strconv.ParseFloat(*r.Time4, 64); err == nil {
			times = append(times, time4)
		}
	}
	sort.Float64s(times)
	return times[0]

}
