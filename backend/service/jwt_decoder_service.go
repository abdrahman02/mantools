package service

import (
	"backend/entities/jwtdecoder"
	"fmt"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

type decodeJWTResponse struct {
	Header map[string]any
	Payload jwt.MapClaims
}

type JWTDecoderService interface {
	DecodeJWT(req *jwtdecoder.FormatRequest) (decodeJWTResponse, error)
}

type jwtDecoderService struct {}

func NewJWTDecoderService() JWTDecoderService {
	return &jwtDecoderService{}
}

func (s jwtDecoderService) DecodeJWT(req *jwtdecoder.FormatRequest) (decodeJWTResponse, error) {
	token, _, err := new(jwt.Parser).ParseUnverified(req.Input, jwt.MapClaims{})
	if err != nil {
		log.Println("Failed to parse token", err.Error())
		return decodeJWTResponse{}, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		log.Println("Invalid claims: ", token.Claims)
		return decodeJWTResponse{}, fmt.Errorf("Invalid claims: %v", token.Claims)
	}
	return decodeJWTResponse{Header: token.Header, Payload: claims}, nil
}