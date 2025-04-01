package models

type Nomination struct {
	Rank               int
	Name               string
	YearOfBirth        string
	Organization       string
	CompetitionName    string
	BestTime           string
	ParticipationCount string
	Qualified          bool
}

func (t *Nomination) SetRank(rank int) {
	t.Rank = rank
}

func (t *Nomination) GetRank() int {
	return t.Rank
}

func (t Nomination) GetTop() string {
	return t.BestTime
}

func (t Nomination) SetTop(top string) {
	t.BestTime = top
}
