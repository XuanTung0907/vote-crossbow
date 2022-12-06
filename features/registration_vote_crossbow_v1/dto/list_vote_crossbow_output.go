package dto

import "time"

type ListRegistrationVoteCrossbow struct {
	CreatedAt   time.Time `json:"create_at"`
	CreatedBy   uint64    `json:"created_by"`
	UpdatedAt   time.Time `json:"updated_at"`
	UpdatedBy   uint64    `json:"updated_by"`
	DeletedAt   time.Time `json:"deleted_at"`
	ID          uint64    `json:"id"`
	Name        string    `json:"name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	Content     string    `json:"content"`
}
