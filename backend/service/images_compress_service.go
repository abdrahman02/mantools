package service

import (
	"archive/zip"
	"backend/entities/imagescompress"
	"bytes"
	"fmt"
	"strconv"

	"github.com/disintegration/imaging"
)

type ImagesCompressService interface {
	ImagesCompress(req *imagescompress.FormatRequest) (*bytes.Buffer, error)
}

type imagesCompressService struct {}

func NewImagesCompressService() ImagesCompressService {
	return &imagesCompressService{}
}

func (s *imagesCompressService) ImagesCompress(req *imagescompress.FormatRequest) (*bytes.Buffer, error) {
	quality, err := strconv.Atoi(req.Quality)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	for _, file := range req.Files {
		fmt.Println("filename: ", file.Filename)
		src, err := file.Open()
		if err != nil {
			fmt.Println("skipping file: ", file.Filename)
			continue
		}
		defer src.Close()

		img, err := imaging.Decode(src)
		if err != nil {
			fmt.Println("skipping file: ", file.Filename)
			continue
		}

		fw, err := zipWriter.Create(file.Filename)
		if err != nil {
			fmt.Println("skipping file: ", file.Filename)
			continue
		}

		err = imaging.Encode(fw, img, imaging.JPEG, imaging.JPEGQuality(quality))
		if err != nil {
			fmt.Println("skipping file: ", file.Filename)
			continue
		}
	}

	zipWriter.Close()

	return buf, err
}