package expensify

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileExportRequest_Validate(t *testing.T) {
	validConfig := fileExportRequest{
		Type: "file",
		InputSettings: exportInputSettings{
			Type:          "combinedReportData",
			EmployeeEmail: "valid@example.com",
			Filters: inputSettingsFilters{
				StartDate:     "2023-01-01",
				EndDate:       "2023-01-31",
				ApprovedAfter: "2023-01-15",
			},
		},
		Test: "true",
		OutputSettings: fileExportOutPutSettings{
			FileExtension:              "pdf",
			IncludeFullPageReceiptsPdf: true,
		},
	}

	tests := []struct {
		name    string
		config  fileExportRequest
		wantErr bool
		errMsg  string
	}{
		{
			name:    "Valid Config",
			config:  validConfig,
			wantErr: false,
		},
		{
			name: "Invalid Type",
			config: fileExportRequest{
				Type: "invalid",
				InputSettings: exportInputSettings{
					Type: "combinedReportData",
				},
				Test: "true",
			},
			wantErr: true,
			errMsg:  "receivied unexpected Type on the request. Exepcted: 'file', received: invalid",
		},
		{
			name: "Invalid InputSettings Type",
			config: fileExportRequest{
				Type: "file",
				InputSettings: exportInputSettings{
					Type: "invalid",
				},
				Test: "true",
			},
			wantErr: true,
			errMsg:  "invalid Type: 'file', valid types are: '' or 'combinedReportData'",
		},
		{
			name: "Invalid Test value",
			config: fileExportRequest{
				Type: "file",
				Test: "invalid",
			},
			wantErr: true,
			errMsg:  "incorrect value received for Test: 'invalid', expecteed values are: 'true' or 'false'",
		},
		{
			name: "Invalid Email Format",
			config: fileExportRequest{
				Type: "file",
				InputSettings: exportInputSettings{
					Type:          "combinedReportData",
					EmployeeEmail: "invalid-email",
				},
				Test: "true",
			},
			wantErr: true,
			errMsg:  "the value passed for InputSettings.EmployeeEmail: 'invalid-email' is not a valid email format",
		},
		{
			name: "Invalid StartDate Format",
			config: fileExportRequest{
				Type: "file",
				InputSettings: exportInputSettings{
					Type: "combinedReportData",
					Filters: inputSettingsFilters{
						StartDate: "01-01-2023",
					},
				},
				Test: "true",
			},
			wantErr: true,
			errMsg:  "this date: '01-01-2023' cannot parsed.",
		},
		{
			name: "Invalid Output FileExtension",
			config: fileExportRequest{
				Type: "file",
				OutputSettings: fileExportOutPutSettings{
					FileExtension: "invalid",
				},
				InputSettings: exportInputSettings{
					Type: "combinedReportData",
					Filters: inputSettingsFilters{
						StartDate: "2023-01-01",
					},
				},
				Test: "true",
			},
			wantErr: true,
			errMsg:  "invalid output fileExtension received: 'invalid' . Valid extensions are:[pdf json]",
		},
		{
			name: "Invalid PDF Settings",
			config: fileExportRequest{
				Type: "file",
				OutputSettings: fileExportOutPutSettings{
					FileExtension:              "json",
					IncludeFullPageReceiptsPdf: true,
				},
				InputSettings: exportInputSettings{
					Type: "combinedReportData",
					Filters: inputSettingsFilters{
						StartDate: "2023-01-01",
					},
				},
				Test: "true",
			},
			wantErr: true,
			errMsg:  "IncludeFullPageReceiptsPdf can only be set to true, if you are exporting a PDF! you are currently trying to export: 'json'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.config.validate()
			if tt.wantErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
