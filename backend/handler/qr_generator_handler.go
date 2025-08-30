package handler

import (
	"backend/entities/qrgenerator"
	"backend/pkg/helper"
	"backend/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type QRGeneratorHandler struct {
	service service.QRGeneratorService
}

func NewQRGeneratorHandler(s service.QRGeneratorService) *QRGeneratorHandler {
	return &QRGeneratorHandler{service: s}
}

func (h QRGeneratorHandler) GenerateQR(ctx *gin.Context) {
	var req qrgenerator.FormatRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println("Invalid request", err)
		helper.BadRequestResponse(ctx, "Invalid request", err.Error())
		return
	}

	png, err := h.service.GenerateQR(&req)
	if err != nil {
		log.Println("Failed create png file", err)
		helper.BadRequestResponse(ctx, "Failed create png file", err.Error())
		return
	}

	ctx.Header("X-Message", "QR is successfully generated!")
	ctx.Header("Content-Disposition", "attachment; filename=qr_generated.png")
	ctx.Header("Access-Control-Expose-Headers", "Content-Disposition, X-Message")
	ctx.Header("Content-Type", "image/png")
	ctx.Data(http.StatusOK, "image/png", png)
}