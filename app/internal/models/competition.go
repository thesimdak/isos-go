package models

import "time"

type Competition struct {
	ID                 int64     `json:"id"`
	CompetitionName    string    `json:"competitionname"`
	Name               string    `json:"name"`
	Date               time.Time `json:"date"`
	Place              string    `json:"place"`
	Judge              string    `json:"judge"`
	SensorInstallation string    `json:"sensorinstallation"`
	Starter            string    `json:"starter"`
	Type               string    `json:"type"`
}

func (c Competition) FormattedDate() string {
	return c.Date.Format("02.01.2006") // Format as DD.MM.YYYY
}
