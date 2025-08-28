package handler

import (
	"backend/entities/imagescompress"
	"backend/pkg/helper"
	"backend/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ImagesCompressHandler struct {
	service service.ImagesCompressService
}

func NewImagesCompressHandler(s service.ImagesCompressService) *ImagesCompressHandler {
	return &ImagesCompressHandler{service: s}
}

func (h ImagesCompressHandler) ImagesCompress(ctx *gin.Context) {
	var req imagescompress.FormatRequest
	if err := ctx.ShouldBind(&req); err != nil {
		log.Println("Invalid request", err)
		helper.BadRequestResponse(ctx, "Invalid request", err)
		return
	}

 	result, err := h.service.ImagesCompress(&req)
 	if err != nil || result == nil {
		log.Println("Failed to compress your files", err)
 		helper.BadRequestResponse(ctx, "Failed to compress your files", err)
 		return
 	}

 	ctx.Header("X-Message", "Compressed successfully!")
 	ctx.Header("Content-Disposition", "attachment; filename=images_compressed.zip")
	ctx.Header("Access-Control-Expose-Headers", "Content-Disposition, X-Message")
 	ctx.Header("Content-Type", "application/zip")
 	ctx.Data(http.StatusOK, "application/zip", result.Bytes())
}