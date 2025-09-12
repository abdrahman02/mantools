package handler

import (
	"backend/entities/auth"
	"backend/pkg/helper"
	"backend/service"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service service.AuthService
}

func NewAuthHandler(s service.AuthService) *AuthHandler { return &AuthHandler{service: s} }

func (h *AuthHandler) Login(ctx *gin.Context) {
	var req auth.FormatRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println("Invalid request", err.Error())
		helper.BadRequestResponse(ctx, "Invalid request", err.Error())
		return
	}

	user, access, refresh, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		log.Println("Failed to process your login request", err)
		helper.BadRequestResponse(ctx, "Failed to process your login request", err.Error())
		return
	}

	secure := os.Getenv("GIN_MODE") == "release"
	ctx.SetCookie("access_token", access, 900, "/", os.Getenv("DOMAIN"), secure, true) // 15 minute
	ctx.SetCookie("refresh_token", refresh, 604_800, "/", os.Getenv("DOMAIN"), secure, true) // 7 days

	data := auth.FormatResponse{Data: map[string]interface{}{
		"user": map[string]interface{}{"id": user.ID, "email": user.Email, "name": user.Name},
		"access": access,
	}}
	helper.SuccessResponse(ctx, "Login successfully", data)
}

func (h *AuthHandler) Refresh(ctx *gin.Context) {
	rt, err := ctx.Cookie("refresh_token")
	if err != nil {
		log.Println("No refresh token", err)
		helper.UnauthorizedResponse(ctx, "No refresh token", err.Error())
		return
	}

	user, access, newRefresh, err := h.service.Refresh(rt)
	if err != nil {
		log.Println("Failed to create your new refresh token", err)
		helper.UnauthorizedResponse(ctx, "Failed to create your new refresh token", err.Error())
		return
	}

	secure := os.Getenv("GIN_MODE") == "release"
	ctx.SetCookie("access_token", access, 900, "/", os.Getenv("DOMAIN"), secure, true)
	ctx.SetCookie("refresh_token", newRefresh, 604_800, "/", os.Getenv("DOMAIN"), secure, true)
	data := auth.FormatResponse{Data: map[string]interface{}{"id": user.ID, "email": user.Email, "name": user.Name}}
	helper.SuccessResponse(ctx, "Success to create your token", data)
}