package routes

import (
	"backend/handler"
	"backend/service"

	"github.com/gin-gonic/gin"
)

func RegisterImagesConverterRoutes(router *gin.RouterGroup) {
	imageConverterService := service.NewImagesConverterService()
	imageConverterHandler := handler.NewImagesConverterHandler(imageConverterService)
	routerGroup := router.Group("/images-converter")
	{
		routerGroup.POST("", imageConverterHandler.ImagesConvert)
	}
}