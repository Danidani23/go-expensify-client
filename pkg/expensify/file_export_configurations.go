package expensify

import "time"

// FileExportBaseConfig  holds the values the user can set
type FileExportBaseConfig struct {
	FilterByReportId              []string
	FilterByPolicyId              []string
	FilterByStartDate             *time.Time
	FilterByEndDate               *time.Time
	FilterByApprovedAfterDate     *time.Time
	FilterByMarkedAsApprovedTag   *string
	FilterByEmployeeEmail         *string
	FilterByReportState           []string
	LimitNumberOfReportsExported  *int
	outputFileExtension           string
	OutputFileBaseName            *string
	OutputIncludeFullPageReceipts bool
	IsThisAtestCall               bool
}

type OnFinishSendEmailConfig struct {
	Recipients []string `json:"recipients"`
	Message    string   `json:"message"`
}

type OnFinishMarkAsExportedConfig struct {
	Label string `json:"label"`
}

type OnFinishSftpUploadDataConfig struct {
	SftpData struct {
		Host     string `json:"host"`
		Login    string `json:"login"`
		Password string `json:"password"`
		Port     int    `json:"port"`
	} `json:"sftpData"`
}
