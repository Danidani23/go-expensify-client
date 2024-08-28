package expensify

import (
	"context"
	"crypto/sha256"
	"fmt"
	"os"
	"path"
)

type Report struct {
	FileExtension string
	Data          []byte
}

// WriteToDisk writes the report out to the specified folder. the filename will be the hash of the content.
func (er *Report) WriteToDisk(ctx context.Context, pathToFolder string) (fileName string, err error) {
	if er.FileExtension == "" {
		return "", fmt.Errorf("error while writing to disk: the fileExtension of the Report is empty")
	}
	if len(er.Data) == 0 {
		return "", fmt.Errorf("it seems that this ExpensifyExport does not contain data, maybe try to call the DownloadReports method first")
	}

	hash := sha256.Sum256(er.Data)
	hashString := fmt.Sprintf("%x", hash) // Convert hash to a hexadecimal string

	fileName = path.Join(pathToFolder, hashString+"."+er.FileExtension)

	// Write the file to disk
	err = os.WriteFile(fileName, er.Data, 0666)
	if err != nil {
		return "", fmt.Errorf("error while writing this file to disk: '%s'. message: %w", fileName, err)
	}
	return fileName, nil
}
