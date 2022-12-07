package user_vote_crossbow_v1

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"gitlab/vote-crossbow/features/registration_vote_crossbow_v1/model"
	"gitlab/vote-crossbow/features/user_vote_crossbow_v1/dto"
)

type VoteCrossbowService interface {
	CreateVoteCrossbow(c *gin.Context, input dto.CreateUserVoteCrossbowInput) (output dto.CreateUserVoteCrossbowOutput, err error)
	LoginVoteCrossbow(c *gin.Context, input dto.LoginUserVoteCrossbowInput) (output string, err error)
}

func ProvideService(repo VoteCrossbowRepository) VoteCrossbowService {
	return service{repo}
}

type service struct {
	repo VoteCrossbowRepository
}

func (s service) CreateVoteCrossbow(c *gin.Context, input dto.CreateUserVoteCrossbowInput) (output dto.CreateUserVoteCrossbowOutput, err error) {
	var register model.VoteCrossbowModel
	err = s.repo.FirstUserVoteCrossbow(&register, input.Name)
	if register.Name != nil {
		return output, err
	}
	copier.Copy(&register, &input)
	err = s.repo.CreateUserVoteCrossbow(&register)
	if err != nil {
		return output, err
	}

	copier.Copy(&output, &register)

	return output, nil
}

func (s service) LoginVoteCrossbow(c *gin.Context, input dto.LoginUserVoteCrossbowInput) (output string, err error) {
	var register model.VoteCrossbowModel
	err = s.repo.LoginUserVoteCrossbow(&register, input.Name, input.Password)
	if register.ID > 0 {
		return "Success", err
	} else {
		return "Login fail !", err
	}
	return
}
