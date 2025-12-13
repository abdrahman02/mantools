package routes

import (
	"backend/configs"
	"backend/handler"
	middleware "backend/middlewares"
	"backend/repository"
	"backend/service"
	"os"

	"github.com/gin-gonic/gin"
)

func RegisterDashboardRoutes(router *gin.RouterGroup) {
	analyticsRepository := repository.NewAnalyticsRepository(configs.AnalyticsService)
	dashboardService := service.NewDashboardService(analyticsRepository, os.Getenv("GA_PROPERTY_ID"))
	dashboardHandler := handler.NewDashboardHandler(dashboardService)
	routerGroup := router.Group("/dashboard")
	{
		routerGroup.GET("/chart", middleware.AuthMiddleware(), dashboardHandler.Dashboard)
	}
}