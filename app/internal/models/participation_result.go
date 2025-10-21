package models

import (
	"sort"
	"strconv"
)

type ParticipationResult struct {
	Id           int
	Rank         int
	Name         string
	YearOfBirth  string
	Organization string
	Time1        *string
	Time2        *string
	Time3        *string
	Time4        *string
	Top          *string
}

func (a ByTopTime) getTopTimes(index int) []float64 {
	result := a[index]
	times := make([]float64, 0, 4)

	if result.Time1 != nil {
		if time1, err := strconv.ParseFloat(*result.Time1, 64); err == nil {
			times = append(times, time1)
		}
	}
	if result.Time2 != nil {
		if time2, err := strconv.ParseFloat(*result.Time2, 64); err == nil {
			times = append(times, time2)
		}
	}
	if result.Time3 != nil {
		if time3, err := strconv.ParseFloat(*result.Time3, 64); err == nil {
			times = append(times, time3)
		}
	}
	if result.Time4 != nil {
		if time4, err := strconv.ParseFloat(*result.Time4, 64); err == nil {
			times = append(times, time4)
		}
	}
	sort.Float64s(times)
	return times
}

// ByTopTime implements sort.Interface for []ParticipationResult based on the Top time.
type ByTopTime []ParticipationResult

func (a ByTopTime) Len() int      { return len(a) }
func (a ByTopTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTopTime) Less(i, j int) bool {
	topTimesI := a.getTopTimes(i)
	topTimesJ := a.getTopTimes(j)

	for i := 0; i < len(topTimesI); i++ {
		if topTimesI[i] == topTimesJ[i] {
			continue
		} else {
			return topTimesI[i] < topTimesJ[i]
		}
	}
	return true
}

func (result ParticipationResult) GetTopTimes() []float64 {
	times := make([]float64, 0, 4)

	if result.Time1 != nil {
		if time1, err := strconv.ParseFloat(*result.Time1, 64); err == nil {
			times = append(times, time1)
		}
	}
	if result.Time2 != nil {
		if time2, err := strconv.ParseFloat(*result.Time2, 64); err == nil {
			times = append(times, time2)
		}
	}
	if result.Time3 != nil {
		if time3, err := strconv.ParseFloat(*result.Time3, 64); err == nil {
			times = append(times, time3)
		}
	}
	if result.Time4 != nil {
		if time4, err := strconv.ParseFloat(*result.Time4, 64); err == nil {
			times = append(times, time4)
		}
	}
	sort.Float64s(times)
	return times
}
