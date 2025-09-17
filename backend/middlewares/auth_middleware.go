package middleware

import (
	"backend/pkg/helper"
	"backend/pkg/helper/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get access token from cookie or authorization
		token, err := ctx.Cookie("access_token")
		if err != nil {
			authHeader := ctx.GetHeader("Authorization")
			if strings.HasPrefix(authHeader, "Bearer ") {
				token = strings.TrimPrefix(authHeader, "Bearer ")
			}
		}

		if token == "" {
			helper.UnauthorizedResponse(ctx, "Missing token", "Missing access token")
			return
		}

		claims, err := auth.VerifyAccessToken(token)
		if err != nil {
			helper.UnauthorizedResponse(ctx, "Invalid or expired access token", err.Error())
			return
		}

		// Save claims so we can use it if we need
		ctx.Set("userID", claims.UserID)
		ctx.Set("email", claims.Email)

		ctx.Next()
	}
}
