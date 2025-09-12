package models

import "time"

type RefreshToken struct {
	ID string `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID string `json:"user_id" gorm:"type:uuid;index;not null"`
	TokenHash string `json:"token_hash" gorm:"not null"`
	ExpiresAt time.Time `json:"expires_at" gorm:"index;not null"`
	Revoked bool `json:"revoked" gorm:"not null;default:false"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}