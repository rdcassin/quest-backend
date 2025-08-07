package models

import (
	"time"
)

type Users struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Username  string    `gorm:"uniqueIndex" json:"username"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	ImageURL  string    `json:"image_url"`
	ClerkID   string    `gorm:"uniqueIndex" json:"clerk_id"`
	Bio       *string   `json:"bio,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (Users) TableName() string {
	return "user"
}
