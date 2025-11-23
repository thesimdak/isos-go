package models

import "time"

type Competition struct {
	ID                 int64
	CompetitionName    string
	Name               string
	Date               time.Time
	Place              *string
	Judge              *string
	SensorInstallation *string
	Starter            *string
	Type               *string
}

func (c Competition) FormattedDate() string {
	return c.Date.Format("02.01.2006") // Format as DD.MM.YYYY
}
