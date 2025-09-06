package models

import "time"

type RefreshToken struct {
	ID string `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID string `gorm:"type:uuid;index;not null"`
	Token string `gorm:"not null"`
	ExpiresAt time.Time `gorm:"index;not null"`
	Revoked bool `gorm:"not null;default:false"`
	CreatedAt time.Time
}