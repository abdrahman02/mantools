package jwtdecoder

type FormatRequest struct {
	Input string `json:"input" binding:"required"`
}

type FormatResponse struct {
	Data string `json:"data"`
}