package routes

import (
	"backend/handler"
	"backend/service"

	"github.com/gin-gonic/gin"
)

func RegisterJWTDecoderRoutes(router *gin.RouterGroup) {
	jwtDecoderService := service.NewJWTDecoderService()
	jwtDecoderHandler := handler.NewJWTDecoderHandler(jwtDecoderService)
	routerGroup := router.Group("/jwt-decoder")
	{
		routerGroup.POST("", jwtDecoderHandler.DecodeJWT)
	}
}