package registration_vote_crossbow_v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type VoteCrossbowController interface {
	GetRegistrationVoteCrossbow(c *gin.Context)
}

func ProvideController(service VoteCrossbowService) VoteCrossbowController {
	return controller{service}
}

type controller struct {
	service VoteCrossbowService
}

func (r controller) GetRegistrationVoteCrossbow(c *gin.Context) {
	user, err := r.service.GetRegistrationVoteCrossbow(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, user)
}
