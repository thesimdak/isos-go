package upload

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thesimdak/goisos/internal/services/competition"
)

type UploadHandler struct {
	*competition.CompetitionService
}

// NewCompetitionRepository creates a new CompetitionRepository instance
func NewUploadHandler(svc *competition.CompetitionService) *UploadHandler {
	return &UploadHandler{CompetitionService: svc}
}

func (h *UploadHandler) Upload(c *gin.Context) {
	file, _, err := c.Request.FormFile("myfile")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	defer file.Close()

	// Initialize specific repositories
	h.CompetitionService.UploadResults(file)
	c.JSON(http.StatusOK, nil)
}
