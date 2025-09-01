package routes

import (
	"backend/handler"
	"backend/service"

	"github.com/gin-gonic/gin"
)

func RegisterHashGeneratorRoutes(router *gin.Engine) {
	hashGeneratorService := service.NewHashGeneartorService()
	hashGeneratorHandler := handler.NewHashGeneratorHandler(hashGeneratorService)
	routerGroup := router.Group("/hash-generator")
	{
		routerGroup.POST("", hashGeneratorHandler.GenerateHash)
	}
}