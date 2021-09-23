package upload

import (
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"time"
)

func UploadLocal(file *multipart.FileHeader, destFolder string) (string, error) {

	sourceFile, err := file.Open()
	if err != nil {
		return "", err
	}
	defer sourceFile.Close()

	folderPath := "uploads/" + destFolder
	err = os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return "", err
	}

	destinationFile := folderPath + "/" + strconv.FormatInt(time.Now().Unix(), 10) + "-" + file.Filename
	destination, err := os.Create(destinationFile)
	if err != nil {
		return "", err
	}

	defer destination.Close()

	if _, err = io.Copy(destination, sourceFile); err != nil {
		return "", err
	}

	return destinationFile, nil
}
