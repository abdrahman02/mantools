package main

import (
	"backend/configs"
	"backend/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func simpleCors(config configs.Config) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", fmt.Sprintf("%v", config.DomainConfig.FrontendBaseURL))
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
	// load .env
	config := configs.LoadConfig()
	router := gin.Default()
	if err := router.SetTrustedProxies(nil); err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	router.Use(simpleCors(config))

	routes.RegisterTextFormatRoutes(router)
	routes.RegisterTextCaseConvRoutes(router)
	routes.RegisterImagesCompressRoutes(router)
	routes.RegisterImagesConverterRoutes(router)
	routes.RegisterQRGeneratorRoutes(router)
	routes.RegisterApiRequestTesterRoutes(router)
	routes.RegisterJWTDecoderRoutes(router)
	routes.RegisterHashGeneratorRoutes(router)

	port := config.DomainConfig.Port
	log.Printf("🚀 Server running on port %d\n", port)
	if err := router.Run(fmt.Sprintf(":%d", port)); err != nil {
		log.Println("Something went wrong: ", err)
	}
}