package main

import (
	"backend/configs"
	"backend/routes"
	"backend/seeders"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func simpleCors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", os.Getenv("FRONTEND_BASE_URL"))
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET,POST,OPTIONS")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusNoContent)
			return
		}
		ctx.Next()
	}
}

func main() {
	configs.DBConnect()
	configs.DBMigrate()

	seeders.UserSeeder()

	router := gin.Default()
	if err := router.SetTrustedProxies(nil); err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	router.Use(simpleCors())

	routes.RegisterTextFormatRoutes(router)
	routes.RegisterTextCaseConvRoutes(router)
	routes.RegisterImagesCompressRoutes(router)
	routes.RegisterImagesConverterRoutes(router)
	routes.RegisterQRGeneratorRoutes(router)
	routes.RegisterApiRequestTesterRoutes(router)
	routes.RegisterJWTDecoderRoutes(router)
	routes.RegisterHashGeneratorRoutes(router)
	routes.RegisterPDFToolsRoutes(router)
	routes.RegisterAuthRoutes(router)

	port := os.Getenv("PORT")
	log.Printf("🚀 Server running on port %s\n", port)
	if err := router.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Println("Something went wrong: ", err)
	}
}