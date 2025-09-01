package handler

import (
	"backend/entities/hashgenerator"
	"backend/pkg/helper"
	"backend/service"
	"log"

	"github.com/gin-gonic/gin"
)

type HashGeneratorHandler struct {
	service service.HashGeneratorService
}

func NewHashGeneratorHandler(s service.HashGeneratorService) *HashGeneratorHandler {
	return &HashGeneratorHandler{service: s}
}

func (h HashGeneratorHandler) GenerateHash(ctx *gin.Context) {
	var req hashgenerator.FormatRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println("Invalid request", err.Error())
		helper.BadRequestResponse(ctx, "Invalid request", err.Error())
		return
	}

	result, err := h.service.GenerateHash(&req)
	if err != nil {
		log.Println("An error occurred during the hashing process", err.Error())
		helper.InternalServerErrorResponse(ctx, "An error occurred during the hashing process", err.Error())
		return
	}

	data := hashgenerator.FormatResponse{Data: result}
	helper.SuccessResponse(ctx, "Hashing is successfully!", data)
}