package models

type Participation struct {
	ID           int64        `json:"id"`
	RopeClimber  *RopeClimber `json:"roleClimber"`
	Competition  *Competition `json:"competition"`
	Category     *Category    `json:"category"`
	ResultRank   int16        `json:"resultRank"`
	TimeList     []*Time      `json:"timeList"`
	Organization string       `json:"organization"`
}
