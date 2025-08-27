package routes

import (
	"backend/handler"
	"backend/service"

	"github.com/gin-gonic/gin"
)

func RegisterImagesCompressRoutes(router *gin.Engine) {
	imagesCompressService := service.NewImagesCompressService()
	imagesCompressHandler := handler.NewImagesCompressHandler(imagesCompressService)
	routerGroup := router.Group("/images-compressor")
	{
		routerGroup.POST("", imagesCompressHandler.ImagesCompress)
	}
}