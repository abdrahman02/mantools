package handler

import (
	"backend/entities/imagescompress"
	"backend/pkg/helper"
	"backend/service"
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

	req.Quality = ctx.PostForm("quality")
	form, _ := ctx.MultipartForm()
	req.Files = form.File["files"]

 	result, err := h.service.ImagesCompress(&req)
 	if err != nil || result == nil {
 		helper.BadRequestResponse(ctx, "Failed to compress your files", err)
 	}

	ctx.Header("Access-Control-Expose-Headers", "Content-Disposition, X-Message")
 	ctx.Header("Content-Type", "application/zip")
 	ctx.Header("Content-Disposition", "attachment; filename=images_compressed.zip")
 	ctx.Header("X-Message", "Compressed successfully!")
 	ctx.Data(http.StatusOK, "application/zip", result.Bytes())
}