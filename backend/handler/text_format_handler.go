package handler

import (
	"backend/entities/textformat"
	"backend/pkg/helper"
	"backend/service"

	"github.com/gin-gonic/gin"
)

func TextFormatHandler(ctx *gin.Context) {
	var req textformat.FormatRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		helper.BadRequestResponse(ctx, "Invalid json body", err.Error())
		return
	}

	var out string

	switch req.Format{
	case "json":
		out, _ = service.FormatJSON(req.Input)
	case "xml":
		out, _ = service.FormatXML(req.Input)
	default:
		helper.BadRequestResponse(ctx, "Target format is unsupported yet", "")
		return
	}

	data := textformat.FormatResponse{Data: out}
	helper.SuccessResponse(ctx, "Success formatted!", data)
}