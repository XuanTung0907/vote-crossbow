package user_vote_crossbow_v1

import (
	"gitlab/vote-crossbow/pkgs/entities"
	"gorm.io/gorm"
)

type VoteCrossbowRepository interface {
	CreateUserVoteCrossbow(userVoteCrossbow interface{}) (err error)
	FirstUserVoteCrossbow(userVoteCrossbow interface{}, name string) (err error)
	LoginUserVoteCrossbow(userVoteCrossbow interface{}, name string, password string) (err error)
}

func ProvideRepository(repo *gorm.DB) database {
	return database{repo}
}

type database struct {
	db *gorm.DB
}

func (d database) CreateUserVoteCrossbow(userVoteCrossbow interface{}) (err error) {
	err = d.db.Table(entities.UserVoteCrossbow{}.TableName()).Create(userVoteCrossbow).Error
	if err != nil {
		return err
	}
	return err
}

func (d database) FirstUserVoteCrossbow(userVoteCrossbow interface{}, name string) (err error) {
	err = d.db.Table(entities.UserVoteCrossbow{}.TableName()).Where("name=?", name).First(&userVoteCrossbow).Error
	if err != nil {
		return err
	}
	return err
}

func (d database) LoginUserVoteCrossbow(userVoteCrossbow interface{}, name string, password string) (err error) {
	err = d.db.Table(entities.UserVoteCrossbow{}.TableName()).Where("name=? and password=?", name, password).First(&userVoteCrossbow).Error
	if err != nil {
		return err
	}
	return err
}
