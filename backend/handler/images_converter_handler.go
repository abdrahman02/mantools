package handler

import (
	"backend/entities/imagesconverter"
	"backend/pkg/helper"
	"backend/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImageConverterHandler struct {
	service service.ImagesConverterService
}

func NewImagesConverterHandler(s service.ImagesConverterService) *ImageConverterHandler {
	return &ImageConverterHandler{service: s}
}

func (h ImageConverterHandler) ImagesConvert(ctx *gin.Context) {
	var req imagesconverter.FormatRequest
	if err := ctx.ShouldBind(&req); err != nil {
		log.Println("Invalid request", err)
		helper.BadRequestResponse(ctx, "Invalid request", err)
		return
	}

	result, err := h.service.ImagesConvert(&req)
	if result == nil || err != nil {
		log.Println("Failed to convert your files", err)
		helper.BadRequestResponse(ctx, "Failed to convert your files", err)
		return
	}

 	ctx.Header("X-Message", "Compressed successfully!")
 	ctx.Header("Content-Disposition", "attachment; filename=images_converted.zip")
	ctx.Header("Access-Control-Expose-Headers", "Content-Disposition, X-Message")
	ctx.Header("Content-Type", "application/zip")
	ctx.Data(http.StatusOK, "application/zip", result.Bytes())
}