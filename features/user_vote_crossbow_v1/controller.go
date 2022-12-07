package user_vote_crossbow_v1

import (
	"github.com/gin-gonic/gin"
	"gitlab/vote-crossbow/features/user_vote_crossbow_v1/dto"
	"net/http"
)

type VoteCrossbowController interface {
	CreateVoteCrossbow(c *gin.Context)
	LoginVoteCrossbow(c *gin.Context)
}

func ProvideController(service VoteCrossbowService) VoteCrossbowController {
	return controller{service}
}

type controller struct {
	service VoteCrossbowService
}

func (r controller) CreateVoteCrossbow(c *gin.Context) {
	var input dto.CreateUserVoteCrossbowInput
	c.Bind(&input)
	user, err := r.service.CreateVoteCrossbow(c, input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, user)
}

func (r controller) LoginVoteCrossbow(c *gin.Context) {
	var input dto.LoginUserVoteCrossbowInput
	c.Bind(&input)
	user, err := r.service.LoginVoteCrossbow(c, input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, user)
}
