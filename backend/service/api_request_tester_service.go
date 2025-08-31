package service

import (
	"backend/entities/apirequesttester"
	"bytes"
	"io"
	"log"
	"net/http"
)

type ApiRequestTesterService interface {
	ApiRequestTester(req *apirequesttester.FormatRequest) ([]byte, error)
}

type apiRequestTesterService struct{}

func NewApiRequestTesterService() ApiRequestTesterService {
	return &apiRequestTesterService{}
}

func (s apiRequestTesterService) ApiRequestTester(req *apirequesttester.FormatRequest) ([]byte, error) {
	newReq, err := http.NewRequest(req.Method, req.Url, bytes.NewBuffer([]byte(req.Body)))
	if err != nil {
		log.Println("Error while created new request: ", err.Error())
		return nil, err
	}

	for k, v := range req.Headers {
		newReq.Header.Set(k, v)
	}

	client := &http.Client{}
	res, err := client.Do(newReq)
	if err != nil {
		log.Println("Error while send new request: ", err.Error())
		return nil, err
	}
	defer res.Body.Close()

	resBody, err := io.ReadAll(res.Body)

	return resBody, nil
}