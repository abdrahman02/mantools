package service

import (
	"archive/zip"
	"backend/entities/imagesconverter"
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"log"
	"path/filepath"
	"strings"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
)

type ImagesConverterService interface {
	ImagesConvert(req *imagesconverter.FormatRequest) (*bytes.Buffer, error)
}

type imagesConverterService struct {}

func NewImagesConverterService() ImagesConverterService {
	return &imagesConverterService{}
}

func (s imagesConverterService) ImagesConvert(req *imagesconverter.FormatRequest) (*bytes.Buffer, error) {
	const outDir = "images_converted/"

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
		* 1. Get target image ext without "."
		* 2. Split and get file name without dot and ext
		* 3. create new filename
		*/
		ext := strings.ToLower(req.TargetConvert)
		baseFilename := strings.TrimSuffix(file.Filename, filepath.Ext(file.Filename))
		newFilename := baseFilename + "." + ext

		fw, err := zipWriter.Create(outDir + newFilename)
		if err != nil {
			failedFiles = append(failedFiles, file.Filename)
			log.Println(fmt.Sprintf("Failed to make entry space zip for file: %s", file.Filename), err)
			continue
		}

		err = convertImage(ext, fw, img)
		if err != nil {
			failedFiles = append(failedFiles, file.Filename)
			log.Println(fmt.Sprintf("Failed to convert image %s to .%s for file: %s", filepath.Ext(file.Filename), ext, file.Filename), err)
			continue
		}
	}

	if len(failedFiles) > 0 {
		return buf, fmt.Errorf("invalid file(s): %v", strings.Join(failedFiles, ", "))
	}
	return buf, nil
}

func convertImage(ext string, fw io.Writer, img image.Image) error {
	var err error
	switch ext {
	case "jpg", "jpeg":
		err = jpeg.Encode(fw, img, &jpeg.Options{Quality: 90})
	case "png":
		err = png.Encode(fw, img)
	case "webp":
		err = webp.Encode(fw, img, &webp.Options{Lossless: true})
	default:
		err = fmt.Errorf("%s format is not supported yet", ext)
	}
	return err
}