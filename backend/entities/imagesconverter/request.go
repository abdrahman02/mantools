package imagesconverter

import "mime/multipart"

type FormatRequest struct {
    TargetConvert string `form:"targetConvert" binding:"required"`
    Files []*multipart.FileHeader `form:"files" binding:"required"`
}