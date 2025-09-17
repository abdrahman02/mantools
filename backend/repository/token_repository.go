package repository

import (
	"backend/models"
	"time"

	"gorm.io/gorm"
)

type TokenRepository interface {
	SaveRefreshToken(token *models.RefreshToken) error
	FindValidToken(hash string) (*models.RefreshToken, error)
	RevokeToken(hash string) error
}

type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepository{db}
}

func (r tokenRepository) SaveRefreshToken(token *models.RefreshToken) error {
	return r.db.Create(token).Error
}

func (r tokenRepository) FindValidToken(hash string) (*models.RefreshToken, error) {
	var rt models.RefreshToken
	err := r.db.First(&rt, "token_hash = ? AND revoked = false AND expires_at > ?", hash, time.Now()).Error
	if err != nil {
		return nil, err
	}

	return &rt, nil
}

func (r tokenRepository) RevokeToken(hash string) error {
	return r.db.Model(&models.RefreshToken{}).Where("token_hash = ?", hash).Update("revoked", true).Error
}
