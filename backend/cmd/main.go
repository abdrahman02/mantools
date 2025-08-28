package main

import (
	"backend/configs"
	"backend/routes"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func simpleCors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
		ctx.Next()
	}
}

func main() {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(simpleCors())

	// load .env
	config := configs.LoadConfig()

	routes.RegisterTextFormatRoutes(router)
	routes.RegisterTextCaseConvRoutes(router)
	routes.RegisterImagesCompressRoutes(router)
	routes.RegisterImagesConverterRoutes(router)

	port := config.DomainConfig.Port
	fmt.Printf("🚀 Server running on port %d\n", port)
	router.Run(fmt.Sprintf(":%d", port))
}