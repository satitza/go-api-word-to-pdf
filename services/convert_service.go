package services

import (
	"bytes"
	"errors"
	"fmt"
	"go-api-word-to-pdf/common"
	"go-api-word-to-pdf/configuration"
	"os"
	"os/exec"
	"path/filepath"
)

func ConvertWordToPdf(osType common.OperatingSystemType, inputFileFullPath, outputFilePath string) (results []byte, err error) {

	var cmd *exec.Cmd

	config, err := configuration.GetConfig()
	if err != nil {
		return nil, fmt.Errorf(`read configuration error: %s`, err)
	}

	switch osType {
	case common.Windows:
		fmt.Println("Running on Windows")
		cmd = exec.Command(config.LibreofficeConfig.WindowsPath, "--headless", "--convert-to", "pdf", inputFileFullPath, "--outdir", outputFilePath)
	case common.Linux:
		fmt.Println("Running on Linux")
		cmd = exec.Command(config.LibreofficeConfig.LinuxPath, "--headless", "--convert-to", "pdf", inputFileFullPath, "--outdir", outputFilePath)
	case common.MacOS:
		fmt.Println("Running on macOS")
	default:
		fmt.Printf("Running on %s\n", osType)
	}

	if cmd == nil {
		return nil, errors.New(`operating system not support for execute this function`)
	}

	// Capture the output (stdout and stderr)
	var out bytes.Buffer
	var stderr bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = &stderr

	// Run the command
	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("command failed with error: %s", stderr.String())
	}

	fmt.Println("Convert docx to pdf success.")
	fmt.Println(fmt.Sprintf("File pdf %s create success.", outputFilePath))

	// Get the name of the converted PDF file
	pdfFile := filepath.Join(outputFilePath, filepath.Base(inputFileFullPath[:len(inputFileFullPath)-len(filepath.Ext(inputFileFullPath))])+".pdf")

	// Open the generated PDF file and read it as a byte array
	results, err = os.ReadFile(pdfFile)
	if err != nil {
		return nil, fmt.Errorf("error reading the PDF file: %v", err)
	}

	// Delete pdf temp file
	err = os.Remove(pdfFile)
	if err != nil {
		fmt.Printf("Error deleting file: %v\n", err)
	}

	// Return the PDF data as a byte array
	return results, nil

}
