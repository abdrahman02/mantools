package hashgenerator

type FormatRequest struct {
	Input string `json:"input" binding:"required"`
	Algorithm string `json:"algorithm" binding:"required"`
	CostFactor int `json:"costFactor,omitEmpty" `
}

type FormatResponse struct {
	Data string `json:"data"`
}