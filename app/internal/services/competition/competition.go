package competition

import (
	"fmt"
	"log"
	"mime/multipart"
	"strconv"

	"github.com/thesimdak/goisos/internal/loaders"
	"github.com/thesimdak/goisos/internal/models"
	"github.com/thesimdak/goisos/internal/repository/category"
	"github.com/thesimdak/goisos/internal/repository/competition"
	"github.com/thesimdak/goisos/internal/repository/participation"
	"github.com/thesimdak/goisos/internal/repository/ropeclimber"
	"github.com/thesimdak/goisos/internal/repository/time"
	"github.com/xuri/excelize/v2"
)

type CompetitionService struct {
	competitionRepo   *competition.CompetitionRepository
	CategoryRepo      *category.CategoryRepository
	RopeClimberRepo   *ropeclimber.RopeClimberRepository
	ParticipationRepo *participation.ParticipationRepository
	TimeRepo          *time.TimeRepository
}

func (svc *CompetitionService) DeleteCompetition(id int64) {
	svc.TimeRepo.DeleteByCompetitionId(id)
	svc.ParticipationRepo.DeleteByCompetitionId(id)
	svc.competitionRepo.Delete(id)
}

func NewCompetitionService(competitionRepository *competition.CompetitionRepository,
	categoryRepository *category.CategoryRepository,
	ropeClimberRepository *ropeclimber.RopeClimberRepository,
	timeRepository *time.TimeRepository,
	participationRepository *participation.ParticipationRepository,
) *CompetitionService {
	return &CompetitionService{competitionRepo: competitionRepository,
		CategoryRepo:      categoryRepository,
		RopeClimberRepo:   ropeClimberRepository,
		TimeRepo:          timeRepository,
		ParticipationRepo: participationRepository}
}

func (svc *CompetitionService) GetSeasons() []int16 {
	seasons, _ := svc.competitionRepo.GetSeasons()
	return seasons
}

func (svc *CompetitionService) GetCompetitions(season int) []models.Competition {
	competitions, _ := svc.competitionRepo.FindAllCompetitionsBySeason(season)
	return competitions
}

func (s *CompetitionService) CreateCompetition(c *models.Competition) *models.Competition {
	if c.CompetitionName == "" || c.Name == "" {
		panic("competition name and name are required")
	}
	storedCompatition := s.competitionRepo.FindCompetition(c.Name, c.Date)
	if storedCompatition != nil {
		return storedCompatition
	}
	return s.competitionRepo.SaveCompetition(c)
}

func (svc *CompetitionService) UploadResults(file multipart.File) {
	f, err := excelize.OpenReader(file)
	if err != nil {
		log.Panicln("Cannot open file: %w", err)
		panic("Cannot open excel file")
	}
	defer f.Close()
	competition := loaders.LoadCompetition(f)
	competition = svc.CreateCompetition(competition)
	categoryI := svc.CategoryRepo.EnsureCategory(loaders.KatI, "Žáci", 4.5)
	categoryII := svc.CategoryRepo.EnsureCategory(loaders.KatII, "Dorostenci", 4.5)
	categoryIII := svc.CategoryRepo.EnsureCategory(loaders.KatIII, "Muži", 8.0)
	categoryIV := svc.CategoryRepo.EnsureCategory(loaders.KatIV, "Senioři", 8.0)
	categoryV := svc.CategoryRepo.EnsureCategory(loaders.KatV, "Ženy a dorostenky", 4.5)
	categoryVI := svc.CategoryRepo.EnsureCategory(loaders.KatVI, "Žákyně", 4.5)

	svc.SaveResults(competition, categoryI, f)
	svc.SaveResults(competition, categoryII, f)
	svc.SaveResults(competition, categoryIII, f)
	svc.SaveResults(competition, categoryIV, f)
	svc.SaveResults(competition, categoryV, f)
	svc.SaveResults(competition, categoryVI, f)
}

func (svc *CompetitionService) SaveResults(competition *models.Competition, category *models.Category, workbook *excelize.File) {
	rows, _ := workbook.GetRows(category.CategoryKey)
	rowIndex := 1
	for {
		if rowIndex >= len(rows) {
			break
		}

		currentRow := rows[rowIndex]

		if len(currentRow) < 3 {
			break
		}
		ropeClimber := &models.RopeClimber{FirstName: currentRow[1], LastName: currentRow[2], YearOfBirth: stringToInt16(currentRow[3])}

		var times []*models.Time

		if len(currentRow) > 5 {
			time1 := &models.Time{Round: 1, Time: toDecimal(currentRow[5])}
			times = append(times, time1)
		}
		if len(currentRow) > 6 {
			time2 := &models.Time{Round: 2, Time: toDecimal(currentRow[6])}
			times = append(times, time2)
		}
		if len(currentRow) > 7 {
			time3 := &models.Time{Round: 3, Time: toDecimal(currentRow[7])}
			times = append(times, time3)
		}
		if len(currentRow) > 8 {
			time4 := &models.Time{Round: 4, Time: toDecimal(currentRow[8])}
			times = append(times, time4)
		}
		rc := svc.RopeClimberRepo.EnsureRopeClimber(ropeClimber)
		svc.TimeRepo.DeleteByCompetitionIdAndCategoryIdAndRopeClimberId(competition.ID, category.ID, rc.ID)
		p := &models.Participation{RopeClimber: rc, Organization: currentRow[4], TimeList: times, Category: category, Competition: competition}
		svc.ParticipationRepo.DeleteByCompetitionIdAndCategoryIdAndRopeClimberId(p)
		participation := svc.ParticipationRepo.InsertParticipation(p)
		svc.TimeRepo.SaveTimes(participation.ID, times)
		rowIndex++
	}

}

func stringToInt16(value string) int16 {
	num, err := strconv.Atoi(value)
	if err != nil {
		log.Println("Cannot convert value " + value + " into int16")
		panic("Cannot convert value " + value + " into int16")
	}
	return int16(num)
}

func toDecimal(value string) string {
	if value == "x" || value == "X" || value == "-" {
		return "999"
	}
	num, err := strconv.ParseFloat(value, 32)
	if err != nil {
		// TODO fix comment
		log.Println("Cannot convert value " + value + " into int16")
		panic("Cannot convert value " + value + " into int16")
	}

	formattedNum := fmt.Sprintf("%.2f", num)
	return formattedNum

}
