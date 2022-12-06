package registration_vote_crossbow_v1

import (
	"gitlab/vote-crossbow/pkgs/entities"
	"gorm.io/gorm"
)

type VoteCrossbowRepository interface {
	GetRegistrationCoteCrossbow(registrationVoteCrossbow interface{}) (err error)
}

func ProvideRepository(repo *gorm.DB) database {
	return database{repo}
}

type database struct {
	db *gorm.DB
}

func (d database) GetRegistrationCoteCrossbow(registrationVoteCrossbow interface{}) (err error) {
	err = d.db.Table(entities.RegistrationVoteCrossbow{}.TableName()).Find(&registrationVoteCrossbow).Error
	if err != nil {
		return err
	}
	return err
}
