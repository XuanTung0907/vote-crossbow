package registration_vote_crossbow_v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gitlab/vote-crossbow/features/registration_vote_crossbow_v1/dto"
	"gitlab/vote-crossbow/features/registration_vote_crossbow_v1/model"
	"gorm.io/gorm"
	"net/http"
)

type VoteCrossbowService interface {
	GetRegistrationVoteCrossbow(c *gin.Context) (output dto.ListRegistrationVoteCrossbow, err error)
}

func ProvideService(repo VoteCrossbowRepository) VoteCrossbowService {
	return service{repo}
}

type service struct {
	repo VoteCrossbowRepository
}

func (s service) GetRegistrationVoteCrossbow(c *gin.Context) (output dto.ListRegistrationVoteCrossbow, err error) {
	var registration model.VoteCrossbowModel
	err = s.repo.GetRegistrationCoteCrossbow(&registration)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatus(http.StatusNotFound)
			return output, err
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err})
		return output, err
	}

	copier.Copy(&output, &registration)

	return output, nil
}
