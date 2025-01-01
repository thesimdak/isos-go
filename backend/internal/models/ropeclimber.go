package models

type RopeClimber struct {
	ID           int64  `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	YearOfBirth  int16  `json:"yearOfBirth"`
	Organization string `json:"organization"`
}
