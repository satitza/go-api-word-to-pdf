package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go-api-word-to-pdf/common"
	"go-api-word-to-pdf/configuration"
	"go-api-word-to-pdf/services"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

func AppConvertWordToPdfHandler(ctx *gin.Context) {

	// Get the current working directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err)
	}

	config, err := configuration.GetConfig()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Get configuration failed: %s", err.Error()))
		return
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("Upload failed: %s", err.Error()))
		return
	}

	inputFileName := fmt.Sprintf("%s.docx", uuid.New())
	inputTempFilePath := filepath.Join(dir, fmt.Sprintf("%s/%s", config.ConvertConfig.WordTempPath, inputFileName))
	outputTempFilePath := filepath.Join(dir, fmt.Sprintf("%s", config.ConvertConfig.PdfTempPath))

	// Save the file to a specified location (e.g., current directory)
	if err := ctx.SaveUploadedFile(file, inputTempFilePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("Could not save file: %s", err.Error()))
		return
	}

	byteArray, err := services.ConvertWordToPdf(common.OperatingSystemType(runtime.GOOS), inputTempFilePath, outputTempFilePath)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, fmt.Sprintf("Convert file failed: %s", err.Error()))
		return
	}

	// TODO clear temp file
	// Delete word temp file
	err = os.Remove(inputTempFilePath)
	if err != nil {
		fmt.Printf("Error deleting file: %v\n", err)
		return
	}

	ctx.Data(http.StatusOK, "application/pdf", byteArray)
}
