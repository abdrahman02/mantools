package qrgenerator

type FormatRequest struct {
	QRContent string `json:"qrContent" binding:"required"`
}