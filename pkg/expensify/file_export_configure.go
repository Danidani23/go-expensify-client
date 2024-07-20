package expensify

import (
	"fmt"
	"github.com/Danidani23/go-expensify-client/pkg/common"
	"strings"
	"time"
)

type FileExportConfig struct {
	FilterByReportId              *[]string
	FilterByPolicyId              *[]string
	FilterByStartDate             *time.Time
	FilterByEndDate               *time.Time
	FilterByApprovedAfterDate     *time.Time
	FilterByMarkedAsApprovedTag   *string
	FilterByEmployeeEmail         *string
	FilterByReportState           *[]string
	LimitNumberOfReportsExported  *int
	OutputFileExtension           string
	OutputFileBaseName            *string
	OutputIncludeFullPageReceipts bool
	IsThisAtestCall               bool
}

func (c *ExpensifyClient) ConfigureFileExport(config FileExportConfig,
	sendEmailOnFinish *OnFinishSendEmail,
	markAsExportedOnFinish *OnFinishMarkAsExported,
	uploadToSftpOnFinish *OnFinishSftpUploadData,
) error {

	e := fileExportRequest{}

	e.Credentials.PartnerUserID = c.partnerUserID
	e.Credentials.PartnerUserSecret = c.partnerUserSecret

	e.Type = "file"
	e.InputSettings.Type = "combinedReportData" // the API only accepts this value, fails if I pass an empty string
	e.OnReceive.ImmediateResponse = []string{"returnRandomFileName"}
	e.isConfigured = true
	e.OutputSettings.FileExtension = config.OutputFileExtension
	e.OutputSettings.IncludeFullPageReceiptsPdf = config.OutputIncludeFullPageReceipts

	if config.IsThisAtestCall {
		e.Test = "true"
	} else {
		e.Test = "false"
	}

	// --- parsing and adding all parameters

	if config.FilterByReportId != nil {
		e.InputSettings.Filters.ReportIDList = strings.Join(*config.FilterByReportId, ",")
	}

	if config.FilterByPolicyId != nil {
		e.InputSettings.Filters.PolicyIDList = strings.Join(*config.FilterByPolicyId, ",")
	}

	if config.FilterByMarkedAsApprovedTag != nil {
		e.InputSettings.Filters.MarkedAsExported = *config.FilterByMarkedAsApprovedTag
	}

	if config.FilterByEmployeeEmail != nil {
		e.InputSettings.EmployeeEmail = *config.FilterByEmployeeEmail
	}

	if config.LimitNumberOfReportsExported != nil {
		if *config.LimitNumberOfReportsExported < 1 {
			return fmt.Errorf("limitNumberOfReportsExported must equalTo or greaterThan 1, we received: '%d'", *config.LimitNumberOfReportsExported)
		}
		e.InputSettings.Limit = fmt.Sprintf("%d", *config.LimitNumberOfReportsExported)
	}

	if config.OutputFileBaseName != nil {
		e.OutputSettings.FileBaseName = *config.OutputFileBaseName
	}

	// DATES
	if config.FilterByStartDate != nil {
		e.InputSettings.Filters.StartDate = config.FilterByStartDate.Format(timeFormat)
	}
	if config.FilterByEndDate != nil {
		e.InputSettings.Filters.EndDate = config.FilterByEndDate.Format(timeFormat)
	}
	if config.FilterByApprovedAfterDate != nil {
		e.InputSettings.Filters.ApprovedAfter = config.FilterByApprovedAfterDate.Format(timeFormat)
	}

	// REPORT STATE
	if config.FilterByReportState != nil {
		// we ensure there are no duplicates
		myStrings := common.RemoveDuplicates(*config.FilterByReportState)

		// we ensure that all passed states are valid
		for _, state := range myStrings {
			if !common.Contains(validReportStates, state) {
				return fmt.Errorf("invalid reportState received as filter: '%s' valid states are: %s", state, validReportStates)
			}
		}

		e.InputSettings.ReportState = strings.Join(myStrings, ",")
	}

	err := e.validate()
	if err != nil {
		return err
	}

	// ON FINISH CONFIG
	if sendEmailOnFinish != nil {
		myOutput := map[string]interface{}{
			"actionName": "email",
			"recipients": strings.Join(sendEmailOnFinish.Recipients, ","),
			"message":    sendEmailOnFinish.Message,
		}
		e.OnFinish = append(e.OnFinish, myOutput)

	}
	if markAsExportedOnFinish != nil {
		myOutput := map[string]interface{}{
			"actionName": "markAsExported",
			"label":      markAsExportedOnFinish.Label,
		}
		e.OnFinish = append(e.OnFinish, myOutput)

	}
	if uploadToSftpOnFinish != nil {
		myOutput := map[string]interface{}{
			"actionName": "sftpUpload",
			"sftpData": map[string]interface{}{
				"host":     uploadToSftpOnFinish.SftpData.Host,
				"login":    uploadToSftpOnFinish.SftpData.Login,
				"password": uploadToSftpOnFinish.SftpData.Password,
				"port":     uploadToSftpOnFinish.SftpData.Port,
			},
		}
		e.OnFinish = append(e.OnFinish, myOutput)
	}

	// -- committing back to the Client

	c.fileExportConfig = &e

	/*
		output, _ := json.MarshalIndent(e, "", "    ")
		fmt.Println(string(output))
	*/
	return nil
}
