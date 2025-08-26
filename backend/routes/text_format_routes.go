package routes

import (
	"backend/handler"
	"backend/service"

	"github.com/gin-gonic/gin"
)

func RegisterTextFormatRoutes(router *gin.Engine) {
	textFormatterService := service.NewTextFormatterService()
	textFormatterHandler := handler.NewTextFormatterHandler(textFormatterService)
	routerGroup := router.Group("/text-formatter")
	{
		routerGroup.POST("", textFormatterHandler.TextFormatter)
	}
}