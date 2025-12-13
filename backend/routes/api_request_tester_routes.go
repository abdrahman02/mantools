package routes

import (
	"backend/handler"
	"backend/service"

	"github.com/gin-gonic/gin"
)

func RegisterApiRequestTesterRoutes(router *gin.RouterGroup) {
	apiRequestTesterService := service.NewApiRequestTesterService()
	apiRequestTesterHandler := handler.NewApiRequestTesterHandler(apiRequestTesterService)
	routerGroup := router.Group("/api-request-tester")
	{
		routerGroup.POST("", apiRequestTesterHandler.ApiRequestTester)
	}
}