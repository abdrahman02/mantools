package service

import (
	"backend/entities/textcaseconv"
	"errors"
	"fmt"
	"strings"
)

type TextCaseConvService interface {
	TextCaseConv(req *textcaseconv.FormatRequest) (string, error)
}

type textCaseConvService struct {}

func NewTextCaseConvService() TextCaseConvService {
	return &textCaseConvService{}
}

func (s *textCaseConvService) TextCaseConv(req *textcaseconv.FormatRequest) (string, error) {
	input := strings.TrimSpace(req.Input)
	var formatted string
	switch req.Format {
	case "uppercase":
		formatted = strings.ToUpper(input)
	case "lowercase":
		formatted = strings.ToLower(input)
	case "capitalize":
		formatted = strings.Title(input)
	default:
		formatted = req.Input
		return formatted, errors.New(fmt.Sprintf("%s format is unsupported yet", req.Format))
	}
	return formatted, nil
}