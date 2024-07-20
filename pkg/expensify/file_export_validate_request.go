package expensify

import (
	"fmt"
	"github.com/Danidani23/go-expensify-client/pkg/common"
	"time"
)

// validates if the settings are correctly filled
func (e *fileExportRequest) validate() error {
	if e.Type != "file" {
		return fmt.Errorf("receivied unexpected Type on the request. Exepcted: 'file', received: %s", e.Type)
	}

	if e.InputSettings.Type != "" && e.InputSettings.Type != "combinedReportData" {
		return fmt.Errorf("invalid Type: '%s', valid types are: '' or 'combinedReportData'", e.Type)
	}

	if e.Test != "true" && e.Test != "false" {
		return fmt.Errorf("incorrect value received for Test: '%s', expecteed values are: 'true' or 'false'", e.Test)
	}

	if e.InputSettings.EmployeeEmail != "" && !common.IsValidEmail(e.InputSettings.EmployeeEmail) {
		return fmt.Errorf("the value passed for InputSettings.EmployeeEmail: '%s' is not a valid email format", e.InputSettings.EmployeeEmail)
	}

	// Validating dates
	for _, date := range []string{
		e.InputSettings.Filters.StartDate,
		e.InputSettings.Filters.EndDate,
		e.InputSettings.Filters.ApprovedAfter,
	} {
		if date != "" {
			_, err := time.Parse("2006-01-02", date)
			if err != nil {
				return fmt.Errorf("this date: '%s' cannot parsed. %w", date, err)
			}
		}
	}

	// Validate report states filters
	var result []string
	common.Combinate(validReportStates, []string{}, &result)

	if !common.Contains(result, e.InputSettings.ReportState) && e.InputSettings.ReportState != "" {
		return fmt.Errorf("this report state: %s is invalid. valid states are:\n%v", e.InputSettings.ReportState, result)
	}

	// Validate mandatory filter combinations
	if e.InputSettings.Filters.ReportIDList == "" && e.InputSettings.Filters.StartDate == "" &&
		e.InputSettings.Filters.EndDate == "" {
		return fmt.Errorf("incorrect filter settings. ReportIDList, StartDate, EndDate cannot be all empty. " +
			"PLease specify at least one of them")
	}

	// output settings - check mandatory
	if !common.Contains(validFileExtensions, e.OutputSettings.FileExtension) {
		return fmt.Errorf("invalid output fileExtension received: '%s' . Valid extensions are:\n%s",
			e.OutputSettings.FileExtension, validFileExtensions)
	}

	// PDF settings
	if e.OutputSettings.FileExtension != "pdf" && e.OutputSettings.IncludeFullPageReceiptsPdf == true {
		return fmt.Errorf("IncludeFullPageReceiptsPdf can only be set to true, if you are exporting a PDF! "+
			"you are currently trying to export: '%s'", e.OutputSettings.FileExtension)
	}

	return nil
}
