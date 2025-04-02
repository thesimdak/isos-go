package models

type RankedResult interface {
	SetRank(rank int)
	GetRank() int
	GetTop() string
	SetTop(top string)
}

type TopParticipationResults struct {
	Id              int
	Rank            int
	Name            string
	YearOfBirth     string
	Organization    string
	CompetitionName string
	Top             string
}

func (t TopParticipationResults) GetRank() int {
	return t.Rank
}

func (t *TopParticipationResults) SetRank(rank int) {
	t.Rank = rank
}

func (t *TopParticipationResults) SetTop(top string) {
	t.Top = top
}

func (t TopParticipationResults) GetTop() string {
	return t.Top
}
