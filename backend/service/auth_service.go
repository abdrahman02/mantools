package service

import (
	"backend/models"
	"backend/pkg/helper/auth"
	"backend/repository"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"

	"gorm.io/gorm"
)

type AuthService interface {
	Login(email, password string) (*models.User, string, string, error)
	Logout(refreshToken string) error
	Refresh(refreshToken string) (*models.User, string, string, error)
}

type authService struct {
	userRepo repository.UserRepository
	tokenRepo repository.TokenRepository
}

func NewAuthService(u repository.UserRepository, t repository.TokenRepository) AuthService {
	return &authService{userRepo: u, tokenRepo: t}
}

func hashToken(token string) string {
	h := sha256.Sum256([]byte(token))
	return hex.EncodeToString(h[:])
}

func (s *authService) Login(email, password string) (*models.User, string, string, error) {
	user, err := s.userRepo.FindByField("email", email)
	if err != nil {
		log.Println("Email or password is wrong: ", err)
		return nil, "", "", err
	}

	if err := auth.CheckPassword(user.Password, password); err != nil {
		log.Println("Email or password is wrong", err)
		return nil, "", "", err
	}

	accessToken, _, err := auth.CreateAccessToken(user.ID, user.Email)
	if err != nil {
		log.Println("Failed to create token", err)
		return nil, "", "", err
	}

	refreshToken, expRefresh, err := auth.CreateRefreshToken(user.ID, user.Email)
	if err != nil {
		log.Println("Failed to create token", err)
		return nil, "", "", err
	}

	rt := &models.RefreshToken{
		UserID: user.ID,
		TokenHash: hashToken(refreshToken),
		ExpiresAt: expRefresh,
	}
	if err := s.tokenRepo.SaveRefreshToken(rt); err != nil { return nil, "", "", err }

	return user, accessToken, refreshToken, nil
}

func (s *authService) Logout(refreshToken string) error {
  if refreshToken == "" { return fmt.Errorf("Missing token") }

  hashed := hashToken(refreshToken)
  return s.tokenRepo.RevokeToken(hashed)
}

func (s *authService) Refresh(refreshToken string) (*models.User, string, string, error) {
	if refreshToken == "" { return nil, "", "", fmt.Errorf("Missing token") }

	_, err := auth.VerifyRefreshToken(refreshToken)
	if err != nil { return nil, "", "", err }

	hashed := hashToken(refreshToken)
	rt, err := s.tokenRepo.FindValidToken(hashed)
	if err != nil {
		if err == gorm.ErrRecordNotFound { return nil, "", "", fmt.Errorf("Refresh token not found") }
		return nil, "", "", err
	}

	user, err := s.userRepo.FindByField("id", rt.UserID)
	if err != nil { return nil, "", "", err }

	// refresh token rotation: revoke the oldest
	err = s.tokenRepo.RevokeToken(hashed)
	if err != nil { return nil, "", "", err }

	// create new token
	accessToken, _, err := auth.CreateAccessToken(user.ID, user.Email)
	if err != nil { return nil, "", "", err }

	newRefresh, expRefresh, err := auth.CreateRefreshToken(user.ID, user.Email)
	if err != nil { return nil, "", "", err }

	rtNew := &models.RefreshToken{UserID: user.ID, TokenHash: hashToken(newRefresh), ExpiresAt: expRefresh}
	if err := s.tokenRepo.SaveRefreshToken(rtNew); err != nil { return nil, "", "", err }

	return user, accessToken, newRefresh, nil
}