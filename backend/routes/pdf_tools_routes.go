package routes

import (
    "backend/handler"
    "backend/service"

    "github.com/gin-gonic/gin"
)

func RegisterPDFToolsRoutes(router *gin.RouterGroup) {
    pdfToolsService := service.NewPDFToolsService()
    pdfToolsHandler := handler.NewPDFToolsHandler(pdfToolsService)
    routerGroup := router.Group("/pdf-tools")
    {
        routerGroup.POST("", pdfToolsHandler.ProcessPDF)
    }
}