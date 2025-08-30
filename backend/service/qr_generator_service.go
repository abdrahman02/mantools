package service

import (
	"backend/entities/qrgenerator"
	"log"

	qrcode "github.com/skip2/go-qrcode"
)

type QRGeneratorService interface {
	GenerateQR(req *qrgenerator.FormatRequest) ([]byte, error)
}

type qrGeneratorService struct {}

func NewQRGeneratorService() QRGeneratorService {
	return &qrGeneratorService{}
}

func (s qrGeneratorService) GenerateQR(req *qrgenerator.FormatRequest) ([]byte, error) {
	const size = 256
	png, err := qrcode.Encode(req.QRContent, qrcode.Medium, size)
	if err != nil {
		log.Println("Failed create png file", err)
		return png, err
	}

	return png, err
}