package handler

import (
	"backend/pkg/helper"
	"backend/service"
	"log"

	"github.com/gin-gonic/gin"
)

type DashboardHandler struct {
	service service.DashboardService
}

func NewDashboardHandler(s service.DashboardService) *DashboardHandler {
	return &DashboardHandler{service: s}
}

func (h *DashboardHandler) Dashboard(ctx *gin.Context) {
	context := ctx.Request.Context()

	data, err := h.service.GetDashboardData(context)
	if err != nil {
		log.Println("Failed to get dashboard data", err)
		helper.BadRequestResponse(ctx, "Failed to get dashboard data", err.Error())
		return
	}

	helper.SuccessResponse(ctx, "Success to get dashboard data", data)
}