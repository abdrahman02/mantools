package handler

import (
	"backend/entities/pdftools"
	"backend/pkg/helper"
	"backend/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PDFToolsHandler struct {
	service service.PDFToolsService
}

func NewPDFToolsHandler(s service.PDFToolsService) *PDFToolsHandler {
	return &PDFToolsHandler{service: s}
}

func (h PDFToolsHandler) ProcessPDF(ctx *gin.Context) {
	var req pdftools.FormatRequest
	if err := ctx.ShouldBind(&req); err != nil {
		log.Println("Invalid request", err.Error())
		helper.BadRequestResponse(ctx, "Invalid request", err.Error())
		return
	}

	result, err := h.service.ProcessPDF(&req)
	if err != nil {
		log.Println("Failed to process your file(s)", err.Error())
		helper.BadRequestResponse(ctx, "Failed to process your file(s)", err.Error())
		return
	}

	type resPackage struct {
		outFilename string
		contentType string
	}

	resp := map[string]resPackage{
		"merge_pdf": resPackage{outFilename: "pdf_merged.pdf", contentType: "application/pdf"},
		"split_pdf": resPackage{outFilename: "pdf_splitted.zip", contentType: "application/zip"},
		"compress_pdf": resPackage{outFilename: "pdf_compressed.zip", contentType: "applicaiton/zip"},
	}

	ctx.Header("X-Message", "Merging is successfully")
	ctx.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", resp[req.PDFAction].outFilename))
	ctx.Header("Access-Control-Expose-Headers", "Content-Disposition, X-Message")
	ctx.Header("Content-Type", resp[req.PDFAction].contentType)
	ctx.Data(http.StatusOK, resp[req.PDFAction].contentType, result.Bytes())
}