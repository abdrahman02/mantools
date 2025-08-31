package handler

import (
	"backend/entities/apirequesttester"
	"backend/pkg/helper"
	"backend/service"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type ApiRequestTesterHandler struct {
	service service.ApiRequestTesterService
}

func NewApiRequestTesterHandler(s service.ApiRequestTesterService) *ApiRequestTesterHandler {
	return &ApiRequestTesterHandler{service: s}
}

func (h ApiRequestTesterHandler) ApiRequestTester(ctx *gin.Context) {
	var req apirequesttester.FormatRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		log.Println("Invalid request: ", err.Error())
		helper.BadRequestResponse(ctx, "Invalid Request", err.Error())
		return
	}

	result, err := h.service.ApiRequestTester(&req)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to reach url: %s", req.Url), err.Error())
		helper.InternalServerErrorResponse(ctx, fmt.Sprintf("Failed to reach url: %s", req.Url), err.Error())
		return
	}

	resultFormatted, err := service.FormatJSON(string(result))
	if err != nil {
		log.Println(fmt.Sprintf("Failed to formatted response from url: %s", req.Url), err.Error())
		helper.InternalServerErrorResponse(ctx, fmt.Sprintf("Failed to formatted response from url: %s", req.Url), err.Error())
		return
	}

	data := apirequesttester.FormatResponse{Data: resultFormatted}
	helper.SuccessResponse(ctx, fmt.Sprintf("Success to hit the url: %s", req.Url), data)
}