package routes

import (
	"backend/handler"
	"backend/service"

	"github.com/gin-gonic/gin"
)

func RegisterQRGeneratorRoutes(router *gin.RouterGroup) {
	qrGeneratorService := service.NewQRGeneratorService()
	qrGeneratorHandler := handler.NewQRGeneratorHandler(qrGeneratorService)
	routerGroup := router.Group("/qr-generator")
	{
		routerGroup.POST("", qrGeneratorHandler.GenerateQR)
	}
}