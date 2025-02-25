package loaders

import (
	"fmt"
	"time"

	"github.com/thesimdak/goisos/internal/models"
	"github.com/xuri/excelize/v2"
)

const (
	InfoSheet              = "INFO"
	CompetitionNameCell    = "B1"
	YearNameCell           = "B2"
	DateCell               = "B3"
	PlaceCell              = "B4"
	JudgeCell              = "B5"
	SensorInstallationCell = "B6"
	StarterCell            = "B7"
	TypeCell               = "B8"
	MaxValue               = 999.0
	KatI                   = "KAT_I"
	KatII                  = "KAT_II"
	KatIII                 = "KAT_III"
	KatIV                  = "KAT_IV"
	KatV                   = "KAT_V"
	KatVI                  = "KAT_VI"
)

func LoadCompetition(f *excelize.File) *models.Competition {

	competition := &models.Competition{}

	// Access cell values by cell address
	competition.CompetitionName, _ = f.GetCellValue(InfoSheet, CompetitionNameCell)
	competition.Name, _ = f.GetCellValue(InfoSheet, YearNameCell)

	dateStr, _ := f.GetCellValue(InfoSheet, DateCell)
	if date, err := time.Parse("2.1.2006", dateStr); err == nil {
		competition.Date = date
	}
	_, err := time.Parse("1.2.2006", dateStr)
	fmt.Print(err)
	competition.Place, _ = f.GetCellValue(InfoSheet, PlaceCell)
	competition.Judge, _ = f.GetCellValue(InfoSheet, JudgeCell)
	competition.SensorInstallation, _ = f.GetCellValue(InfoSheet, SensorInstallationCell)
	competition.Starter, _ = f.GetCellValue(InfoSheet, StarterCell)
	competition.Type, _ = f.GetCellValue(InfoSheet, TypeCell)

	return competition
}
