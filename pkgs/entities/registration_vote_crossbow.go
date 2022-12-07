package entities

import "time"

type RegistrationVoteCrossbow struct {
	CreatedAt   time.Time `json:"create_at" gorm:"create_at"`
	CreatedBy   uint64    `json:"created_by" gorm:"created_by"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"updated_at"`
	UpdatedBy   uint64    `json:"updated_by" gorm:"updated_by"`
	DeletedAt   time.Time `json:"deleted_at" gorm:"deleted_at"`
	ID          uint64    `json:"id" gorm:"id"`
	Name        *string   `json:"name" gorm:"name"`
	PhoneNumber string    `json:"phone_number" gorm:"phone_number"`
	Email       string    `json:"email" gorm:"email"`
	Content     string    `json:"content" gorm:"content"`
}

func (RegistrationVoteCrossbow) TableName() string {
	return "registration_vote_crossbow"
}
