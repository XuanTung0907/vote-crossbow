package entities

type UserVoteCrossbow struct {
	Id       int    `json:"id" gorm:"id"`
	Name     string `json:"name" gorm:"name"`
	Company  string `json:"company" gorm:"company"`
	Password string `json:"password" gorm:"password"`
}

func (UserVoteCrossbow) TableName() string {
	return "register_vote_crossbow"
}
