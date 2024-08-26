package expensify

import (
	"context"
	"fmt"
	"os"
	"path"
)

type ExpensifyReport struct {
	FileName   string `json:"fileName"`
	FileSystem string `json:"fileSystem"`
	Data       []byte `json:"data"`
}

type downloaderConfig struct {
	Type        string         `json:"type"`
	Credentials expCredentials `json:"credentials"`
	FileName    string         `json:"fileName"`
	FileSystem  string         `json:"fileSystem"`
}

// --------------------------- METHODS

func (er *ExpensifyReport) validate() error {
	if er.FileSystem != "reconciliation" && er.FileSystem != "integrationServer" {
		return fmt.Errorf("incorrect ExpensifyReport, er.FileSystem must be one of ['reconciliation','integrationServer']")
	}

	if er.FileName == "" {
		return fmt.Errorf("incorrect ExpensifyReport, er.FileName cannot be empty")
	}
	return nil
}

func (er *ExpensifyReport) WriteToDisk(ctx context.Context, filePath string) error {
	err := er.validate()
	if err != nil {
		return err
	}
	if len(er.Data) == 0 {
		return fmt.Errorf("it seems that this ExpensifyReport does not contain data, maybe try to call the DownloadReport method first")
	}

	err = os.WriteFile(path.Join(filePath, er.FileName), er.Data, 0666)
	if err != nil {
		return fmt.Errorf("error while writing  this file to disk: '%s'. message: %w", er.FileName, err)
	}
	return nil
}
