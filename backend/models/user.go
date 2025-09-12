package models

import "time"

type User struct {
	ID string `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email string `json:"email" gorm:"uniqueIndex;not null"`
	Password string `json:"password" gorm:"not null"`
	Name string `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}