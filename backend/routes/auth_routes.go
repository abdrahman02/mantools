package routes

import (
	"backend/configs"
	"backend/handler"
	middleware "backend/middlewares"
	"backend/repository"
	"backend/service"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(router *gin.Engine) {
	userRepo := repository.NewUserRepository(configs.DB)
	tokenRepo := repository.NewTokenRepository(configs.DB)
	authService := service.NewAuthService(userRepo, tokenRepo)
	authHandler := handler.NewAuthHandler(authService)
	routerGroup := router.Group("/auth")
	{
		routerGroup.POST("/login", authHandler.Login)
		routerGroup.POST("/refresh", authHandler.Refresh)
		routerGroup.POST("/logout", middleware.AuthMiddleware(), authHandler.Logout)
	}
}