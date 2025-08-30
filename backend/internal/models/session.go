package models

import (
	"time"

	"github.com/google/uuid"
	// "gorm.io/gorm"
)

type SessionUserData struct {
	Email string
}

type SessionDeviceData struct {
	IpAddress string
}

type Session struct {
	ID         uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey" json:"id"`
	UserID     uuid.UUID      `gorm:"type:uuid;not null;index" json:"user_id"`
	// User       User           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user"`
	ExpiresAt  time.Time      `json:"expires_at"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	LastActive time.Time
}

// func (s *Session) BeforeCreate(tx *gorm.DB) (err error) {
// 	s.ID = uuid.New()
// 	return
// }