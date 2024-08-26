package expensify

import (
	"context"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestDownloaderConfig_Validate(t *testing.T) {
	tests := []struct {
		name    string
		config  ExpensifyReport
		wantErr bool
		errMsg  string
	}{
		{
			name: "Valid Config - Reconciliation",
			config: ExpensifyReport{
				FileName:   "report.csv",
				FileSystem: "reconciliation",
			},
			wantErr: false,
		},
		{
			name: "Valid Config - IntegrationServer",
			config: ExpensifyReport{
				FileName:   "report.csv",
				FileSystem: "integrationServer",
			},
			wantErr: false,
		},
		{
			name: "Invalid FileSystem",
			config: ExpensifyReport{
				FileName:   "report.csv",
				FileSystem: "invalidSystem",
			},
			wantErr: true,
			errMsg:  "incorrect ExpensifyReport, er.FileSystem must be one of ['reconciliation','integrationServer']",
		},
		{
			name: "Empty FileName",
			config: ExpensifyReport{
				FileName:   "",
				FileSystem: "reconciliation",
			},
			wantErr: true,
			errMsg:  "incorrect ExpensifyReport, er.FileName cannot be empty",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.validate()
			if tt.wantErr {
				assert.Error(t, err)
				assert.EqualError(t, err, tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestExpensifyReport_WriteToDisk(t *testing.T) {
	// Create a temporary directory for the test
	tempDir := t.TempDir()

	tests := []struct {
		name      string
		report    ExpensifyReport
		filePath  string
		wantErr   bool
		errMsg    string
		checkFile bool
	}{
		{
			name: "Valid Write to Disk",
			report: ExpensifyReport{
				FileName:   "report.txt",
				FileSystem: "reconciliation",
				Data:       []byte("This is a test report"),
			},
			filePath:  tempDir,
			wantErr:   false,
			checkFile: true,
		},
		{
			name: "No Data in Report",
			report: ExpensifyReport{
				FileName:   "report.txt",
				FileSystem: "reconciliation",
				Data:       []byte(""),
			},
			filePath: tempDir,
			wantErr:  true,
			errMsg:   "it seems that this ExpensifyReport does not contain data, maybe try to call the DownloadReport method first",
		},
		{
			name: "Write File Error",
			report: ExpensifyReport{
				FileName:   "report.txt",
				FileSystem: "reconciliation",
				Data:       []byte("This is a test report"),
			},
			filePath: "/invalid/path",
			wantErr:  true,
			errMsg:   "error while writing  this file to disk: 'report.txt'. message: open /invalid/path/report.txt: no such file or directory",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.report.WriteToDisk(context.Background(), tt.filePath)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			} else {
				assert.NoError(t, err)
				if tt.checkFile {
					filePath := filepath.Join(tt.filePath, tt.report.FileName)
					_, err := os.Stat(filePath)
					assert.NoError(t, err)
					assert.FileExists(t, filePath)
				}
			}
		})
	}
}
