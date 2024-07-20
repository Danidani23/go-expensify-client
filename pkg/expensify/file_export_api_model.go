package expensify

import "encoding/json"

type fileExportRequest struct {
	Type      string `json:"type"`
	OnReceive struct {
		ImmediateResponse []string `json:"immediateResponse"`
	} `json:"onReceive"`
	Credentials    expCredentials           `json:"credentials"`
	InputSettings  exportInputSettings      `json:"inputSettings,omitempty"`
	OutputSettings fileExportOutPutSettings `json:"outputSettings,omitempty"`
	Test           string                   `json:"test"`               // true, false If set to true, actions defined in onFinish will not be executed.
	OnFinish       []map[string]interface{} `json:"onFinish,omitempty"` //Actions performed at the end of the export.
	isConfigured   bool
}

func (e *fileExportRequest) marhsall() ([]byte, string, error) {
	jBytes, err := json.Marshal(e)
	if err != nil {
		return nil, "", err
	}
	return jBytes, string(jBytes), nil
}

type expCredentials struct {
	PartnerUserID     string `json:"partnerUserID"`
	PartnerUserSecret string `json:"partnerUserSecret"`
}

// ---------- INPUT

type exportInputSettings struct {
	Type        string               `json:"type,omitempty"` //Specifies all Expensify reports will be combined into a single file.
	Filters     inputSettingsFilters `json:"filters,omitempty"`
	ReportState string               `json:"reportState,omitempty"`
	//Only the reports matching the specified status(es) will be exported. One or more of "OPEN", "SUBMITTED", "APPROVED", "REIMBURSED", "ARCHIVED".
	//When using multiple statuses, separate them by a comma, e.g. "APPROVED,REIMBURSED"
	//Note: These values respectively match the statuses "Open", "Processing", "Approved", "Reimbursed" and "Closed" on the website

	Limit         string `json:",omitempty"`              // Maximum number of reports to export. Any integer, as a string
	EmployeeEmail string `json:"employeeEmail,omitempty"` // The reports will be exported from that account.
	//A valid email address
	//Note:
	//* The usage of this parameter is restricted to certain domains
	//* If this parameter is used, reports in the OPEN status cannot be exported

}

type inputSettingsFilters struct {
	ReportIDList     string `json:"reportIDList,omitempty"`     //Comma-separated list of report IDs to be exported.
	PolicyIDList     string `json:"policyIDList,omitempty"`     //Comma-separated list of policy IDs the exported reports must be under.
	StartDate        string `json:"startDate,omitempty"`        // Filters out all reports submitted or created before the given date, whichever occurred last (inclusive).
	EndDate          string `json:"endDate,omitempty"`          // Filters out all reports submitted or created after the given date, whichever occurred last (inclusive).
	ApprovedAfter    string `json:"approvedAfter,omitempty"`    // Filters out all reports approved before, or on that date. This filter is only used against reports that have been approved.
	MarkedAsExported string `json:"markedAsExported,omitempty"` // Filters out reports that have already been exported with that label out.
}

// ---------- OUTPUT

type fileExportOutPutSettings struct {
	FileExtension string `json:"fileExtension"` // Specifies the format of the generated report.
	// Note: if the "pdf" option is chosen, one PDF file will be generated for each report exported.
	FileBaseName               string `json:"fileBaseName,omitempty"`               // The name of the generated file(s) will start with this value, and a random part will be added to make each filename globally unique. If not specified, the default value export is used.
	IncludeFullPageReceiptsPdf bool   `json:"includeFullPageReceiptsPdf,omitempty"` // Specifies whether generated PDFs should include full page receipts. This parameter is used only if fileExtension contains pdf
}
