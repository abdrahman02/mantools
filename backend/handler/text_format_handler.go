package handler

import (
	"backend/entities/textformat"
	"backend/pkg/helper"
	"backend/service"

	"github.com/gin-gonic/gin"
)

type TextFormatterHandler struct {
	service service.TextFormatterService
}

func NewTextFormatterHandler(s service.TextFormatterService) *TextFormatterHandler {
	return &TextFormatterHandler{service: s}
}

func (h TextFormatterHandler) TextFormatter(ctx *gin.Context) {
	var req textformat.FormatRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.BadRequestResponse(ctx, "Invalid json body", err.Error())
		return
	}

	result, err := h.service.TextFormatter(&req)
	if err != nil {
		helper.BadRequestResponse(ctx, "Target format is unsupported yet", err)
	}

	data := textformat.FormatResponse{Data: result}
	helper.SuccessResponse(ctx, "Success formatted!", data)
}