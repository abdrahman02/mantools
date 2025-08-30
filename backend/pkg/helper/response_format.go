package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Status int `json:"status"` // HTTP status code
	Message string `json:"message"` // short mesage
	Content interface{} `json:"content"` // data
	Error interface{} `json:"error"` // error detail
}

func SuccessResponse(ctx *gin.Context, message string, data interface{}) {
	ctx.JSON(http.StatusOK, BaseResponse{
		Status: http.StatusOK,
		Message: message,
		Content: data,
	})
}

func BadRequestResponse(ctx *gin.Context, message string, err interface{}) {
	ctx.JSON(http.StatusBadRequest, BaseResponse{
		Status: http.StatusBadRequest,
		Message: message,
		Error: err,
	})
}

func UnauthorizedResponse(ctx *gin.Context, message string, err interface{}) {
	ctx.JSON(http.StatusUnauthorized, BaseResponse{
		Status: http.StatusUnauthorized,
		Message: message,
		Error: err,
	})
}

func NotFoundResponse(ctx *gin.Context, message string, err interface{}) {
	ctx.JSON(http.StatusNotFound, BaseResponse{
		Status: http.StatusNotFound,
		Message: message,
		Error: err,
	})
}

func InternalServerErrorResponse(ctx *gin.Context, message string, err interface{}) {
	ctx.JSON(http.StatusInternalServerError, BaseResponse{
		Status: http.StatusInternalServerError,
		Message: message,
		Error: err,
	})
}