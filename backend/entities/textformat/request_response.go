package textformat

type FormatRequest struct {
	Input string `json:"input"`
	Format string `json:"format"`
}

type FormatResponse struct {
	Data string `json:"data"`
}