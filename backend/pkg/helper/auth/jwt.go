package auth

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	UserID string `json:"uid"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func createToken(userID, email, secret string, ttl time.Duration) (string, time.Time, error) {
	now := time.Now()
	exp := now.Add(ttl)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		UserID: userID,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt: jwt.NewNumericDate(now),
		},
	})
	s, err := token.SignedString([]byte(secret))
	return s, exp, err
}

func CreateAccessToken(userID, email string) (string, time.Time, error) {
	ttlMin, _ := strconv.Atoi(os.Getenv("JWT_ACCESS_TTL_MIN"))
	return createToken(userID, email, os.Getenv("JWT_ACCESS_SECRET"), time.Duration(ttlMin) * time.Minute)
}

func CreateRefreshToken(userID, email string) (string, time.Time, error) {
	ttlH, _ := strconv.Atoi(os.Getenv("JWT_REFRESH_TTL_H"))
	return createToken(userID, email, os.Getenv("JWT_REFRESH_SECRET"), time.Duration(ttlH) * time.Hour)
}

func VerifyRefreshToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (any, error) { return []byte(os.Getenv("JWT_REFRESH_SECRET")), nil })
	if err != nil { return nil, err }

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid { return nil, fmt.Errorf("Invalid refresh token") }

	return claims, nil
}

func VerifyAccessToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) { return []byte(os.Getenv("JWT_ACCESS_SECRET")), nil })
	if err != nil { return nil, err }

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid { return nil, fmt.Errorf("Invalid access token") }

	return claims, nil
}