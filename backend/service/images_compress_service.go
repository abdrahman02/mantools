package service

import (
	"archive/zip"
	"backend/entities/imagescompress"
	"bytes"
	"fmt"
	"log"
	"path/filepath"
	"strconv"
	"strings"

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
	const outDir = "images_compressed/"

	quality, err := strconv.Atoi(req.Quality)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)
	defer zipWriter.Close()

	var failedFiles []string
	for _, file := range req.Files {
		src, err := file.Open()
		if err != nil {
			failedFiles = append(failedFiles, file.Filename)
			log.Println(fmt.Sprintf("Failed to open file: %s", file.Filename), err)
			continue
		}

		img, err := imaging.Decode(src)
		src.Close()
		if err != nil {
			failedFiles = append(failedFiles, file.Filename)
			log.Println(fmt.Sprintf("Failed to convert to object image for file: %s", file.Filename), err)
			continue
		}

		/**
		* 1. Split and get file name without dot and ext
		* 2. create new filename
		*/
		baseFilename := strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename))
		newFilename := baseFilename + ".jpeg"
		fw, err := zipWriter.Create(outDir + newFilename)
		if err != nil {
			failedFiles = append(failedFiles, file.Filename)
			log.Println(fmt.Sprintf("Failed to make entry space zip for file: %s", file.Filename), err)
			continue
		}

		err = imaging.Encode(fw, img, imaging.JPEG, imaging.JPEGQuality(quality))
		if err != nil {
			failedFiles = append(failedFiles, file.Filename)
			log.Println(fmt.Sprintf("Failed to compress image for file: %s", file.Filename), err)
			continue
		}
	}

	if len(failedFiles) > 0 {
		return buf, fmt.Errorf("invalid file(s): %v", strings.Join(failedFiles, ", "))
	}

	return buf, err
}