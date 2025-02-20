package models

type Participation struct {
	ID           int64
	RopeClimber  *RopeClimber
	Competition  *Competition
	Category     *Category
	ResultRank   int16
	TimeList     []*Time
	Organization string
}
