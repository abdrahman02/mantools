package routes

import (
	"backend/handler"
	"backend/service"

	"github.com/gin-gonic/gin"
)

func RegisterTextCaseConvRoutes(router *gin.Engine) {
	textCaseConvService := service.NewTextCaseConvService()
	textCaseConvHandler := handler.NewTextCaseConvHandler(textCaseConvService)

	routerGroup := router.Group("/text-case-converter")
	{
		routerGroup.POST("", textCaseConvHandler.TextCaseConv)
	}
}