package handler

import (
	"backend/entities/jwtdecoder"
	"backend/pkg/helper"
	"backend/service"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

type JWTDecoderHandler struct {
	service service.JWTDecoderService
}

func NewJWTDecoderHandler(s service.JWTDecoderService) *JWTDecoderHandler {
	return &JWTDecoderHandler{service: s}
}

func (h JWTDecoderHandler) DecodeJWT(ctx *gin.Context) {
	var req jwtdecoder.FormatRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println("Invalid request", err.Error())
		helper.BadRequestResponse(ctx, "Invalid request", err.Error())
		return
	}

	parts := strings.Split(req.Input, ".")
	if len(parts) != 3 {
		log.Println("Invalid jwt format")
		helper.BadRequestResponse(ctx, "Invalid jwt format", fmt.Errorf("Invalid jwt format"))
		return
	}

	result, err := h.service.DecodeJWT(&req)
	if err != nil {
		log.Println("Failed decode your json web token", err.Error())
		helper.InternalServerErrorResponse(ctx, "Failed decode your json web token", err.Error())
		return
	}

	resultStr, err := json.Marshal(result)
	if err != nil {
		log.Println("Failed convert struct service.decodeJWTResponse to json string", err.Error())
		helper.InternalServerErrorResponse(ctx, "Failed convert struct service.decodeJWTResponse to json string", err.Error())
		return
	}

	resultFormatted, err := service.FormatJSON(string(resultStr))
	if err != nil {
		log.Println("Failed to format result", err.Error())
		helper.InternalServerErrorResponse(ctx, "Failed to format result", err.Error())
		return
	}
	data := jwtdecoder.FormatResponse{Data: resultFormatted}
	helper.SuccessResponse(ctx, "JWT decoded successfully", data)
}