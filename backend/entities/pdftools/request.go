package pdftools

import "mime/multipart"

type FormatRequest struct {
	PDFAction string `form:"pdfAction" binding:"required"`
	Files []*multipart.FileHeader `form:"files" binding:"required"`
	RangePage string `form:"rangePage"`
}