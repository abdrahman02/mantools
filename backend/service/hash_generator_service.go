package service

import (
	"backend/entities/hashgenerator"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type HashGeneratorService interface {
	GenerateHash(req *hashgenerator.FormatRequest) (string, error)
}

type hashGeneratorService struct{}

func NewHashGeneartorService() HashGeneratorService {
	return &hashGeneratorService{}
}

func (hashGeneratorService) GenerateHash(req *hashgenerator.FormatRequest) (string, error) {
	switch req.Algorithm {
	case "md5":
		out := md5.Sum([]byte(req.Input))
		return hex.EncodeToString(out[:]), nil
	case "sha1":
		out := sha1.Sum([]byte(req.Input))
		return hex.EncodeToString(out[:]), nil
	case "sha256":
		out := sha256.Sum256([]byte(req.Input))
		return hex.EncodeToString(out[:]), nil
	case "bcrypt":
		if req.CostFactor < 4 { req.CostFactor = 10 }
		if req.CostFactor > 16 { req.CostFactor = 16 }
		out, err := bcrypt.GenerateFromPassword([]byte(req.Input), req.CostFactor)
		if err != nil {
			log.Println("Hash bcrypt is failed!", err.Error())
			return "", err
		}
		return string(out), nil
	default:
		return "", fmt.Errorf("Unknown algorithm")
	}
}