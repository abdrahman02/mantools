package imagescompress

import "mime/multipart"

type FormatRequest struct {
	Quality string `form:"quality"`
	Files []*multipart.FileHeader `form:"files"`
}