package handler

import (
	"backend/entities/textcaseconv"
	"backend/pkg/helper"
	"backend/service"

	"github.com/gin-gonic/gin"
)

type TextCaseConvHandler struct {
	service service.TextCaseConvService
}

func NewTextCaseConvHandler(s service.TextCaseConvService) *TextCaseConvHandler {
	return &TextCaseConvHandler{service: s}
}

func (h TextCaseConvHandler) TextCaseConv(ctx *gin.Context) {
	var req textcaseconv.FormatRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.BadRequestResponse(ctx, "Invalid json body", err.Error())
		return
	}

	result, err := h.service.TextCaseConv(&req)
	if err != nil {
		helper.BadRequestResponse(ctx, "Target format is unsupported yet", err)
		return
	}

	data := textcaseconv.FormatResponse{Data: result}
	helper.SuccessResponse(ctx, "Success formatted!", data)
}