package apirequesttester

type FormatRequest struct {
	Method string `json:"method" binding:"required"`
	Url string `json:"url" binding:"required"`
	Headers map[string]string `json:"headers" binding:"required"`
	Body string `json:"body"`
}

type FormatResponse struct {
	Data string `json:"data"`
}