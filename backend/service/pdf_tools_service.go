package service

import (
	"archive/zip"
	"backend/entities/pdftools"
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

type PDFToolsService interface {
	ProcessPDF(req *pdftools.FormatRequest) (*bytes.Buffer, error)
}

type pdfToolsService struct {}

func NewPDFToolsService() PDFToolsService {
	return &pdfToolsService{}
}

func (pdfToolsService) ProcessPDF(req *pdftools.FormatRequest) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)

	switch req.PDFAction {
	case "merge_pdf":
		err := mergePDF(buf, req.Files)
		if err != nil {
			return nil, err
		}
	case "split_pdf":
		err := splitPDF(buf, req.Files[0], req.RangePage)
		if err != nil {
			return nil, err
		}
	case "compress_pdf":
		err := compressPDF(buf, req.Files)
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("Target pdf action is not supported yet")
	}
	return buf, nil
}

func mergePDF(buf *bytes.Buffer, files []*multipart.FileHeader) error {
	// create temporary file(s)
	var inFiles []string
	for _, file := range files {
		tempFile, err := saveTempFile(file)
		if err != nil {
			log.Println("Failed to create temporary file", err)
			return err
		}
		defer os.Remove(tempFile)
		inFiles = append(inFiles, tempFile)
	}

	err := pdfcpu.Merge("", inFiles, buf, nil, false)
	if err != nil {
		log.Println("Failed to merging files", err)
		return err
	}

	return nil
}

func splitPDF(buf *bytes.Buffer, file *multipart.FileHeader, rangePage string) error {
	zipWriter := zip.NewWriter(buf)
	defer zipWriter.Close()

	// create temporary file
	tempFile, err := saveTempFile(file)
	if err != nil {
		log.Println("Failed to create temporary file", err)
		return err
	}
	defer os.Remove(tempFile)

	outDir, err := os.MkdirTemp("", "pdf_extract_*")
	if err != nil {
		log.Println("Failed to create temporary directory", err)
		return err
	}
	defer os.RemoveAll(outDir)

	f, err := os.Open(tempFile)
	if err != nil {
		log.Println("Failed to open temp file", err)
		return err
	}
	defer f.Close()

	if err := pdfcpu.ExtractPages(f, outDir, "extracted", []string{rangePage}, nil); err != nil {
		log.Println("Failed to extract", err)
		return err
    }

    err = filepath.Walk(outDir, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        if !info.Mode().IsRegular() {
            return nil
        }

        relPath, err := filepath.Rel(outDir, path)
        if err != nil {
            return err
        }

        writer, err := zipWriter.Create(relPath)
        if err != nil {
            return err
        }

        inFile, err := os.Open(path)
        if err != nil {
            return err
        }
        defer inFile.Close()

        _, err = io.Copy(writer, inFile)
        return err
    })
    if err != nil {
    	log.Println("Failed to zip files", err)
		return err
    }

    return nil
}

func compressPDF(buf *bytes.Buffer, files []*multipart.FileHeader) error {
	zipWriter := zip.NewWriter(buf)
	defer zipWriter.Close()

	for _, file := range files {
		// create temporary file
		tempFile, err := saveTempFile(file)
		if err != nil {
			log.Println("Failed to create temporary file", err)
			return err
		}
		defer os.Remove(tempFile)

		// create temporary file for compress output later
		outFile, err := os.CreateTemp("", "compressed_*.pdf")
		if err != nil {
			log.Println("Failed to create temp file", err)
			return err
		}

		outFilePath := outFile.Name()
		outFile.Close()
		defer os.Remove(outFilePath)

		if err = pdfcpu.OptimizeFile(tempFile, outFilePath, nil); err != nil {
			log.Println("Failed to compress pdf", err)
			return err
		}

		writer, err := zipWriter.Create(file.Filename)
		if err != nil {
			log.Println("Failed to create zip entry", err)
			return err
		}

		compressedFile, err := os.Open(outFilePath)
		if err != nil {
			log.Println("Failed to open compressed file", err)
			return err
		}
		defer compressedFile.Close()

		if _, err := io.Copy(writer, compressedFile); err != nil {
			log.Println("Failed to write to zip", err)
			return err
		}
	}
	return nil
}

func saveTempFile(fileHeader *multipart.FileHeader) (string, error) {
	src, err := fileHeader.Open()
	if err != nil {
		log.Println(fmt.Sprintf("Failed to open file: %s", fileHeader.Filename), err)
		return "", err
	}
	defer src.Close()

	tmpFile, err := os.CreateTemp("", "*.pdf")
	if err != nil {
		log.Println(fmt.Sprintf("Failed to create temporary file: %s", fileHeader.Filename), err)
		return "", err
	}
	defer tmpFile.Close()

	_, err = io.Copy(tmpFile, src)
	if err != nil {
		log.Println(fmt.Sprintf("Failed to copy file %s to temporary file %s", fileHeader.Filename, tmpFile.Name()), err)
		return "", err
	}

	return tmpFile.Name(), nil
}