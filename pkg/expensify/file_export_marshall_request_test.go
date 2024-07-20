package expensify

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExportMarshal(t *testing.T) {

	myReq := fileExportRequest{
		Type: "file",
		Credentials: expCredentials{
			PartnerUserID:     "myId",
			PartnerUserSecret: "myPassword",
		},
		InputSettings: exportInputSettings{
			Type: "",
			Filters: inputSettingsFilters{
				ReportIDList:     "",
				PolicyIDList:     "",
				StartDate:        "2023-01-01",
				EndDate:          "",
				ApprovedAfter:    "",
				MarkedAsExported: "",
			},
			ReportState:   "",
			Limit:         "100",
			EmployeeEmail: "",
		},
		OutputSettings: fileExportOutPutSettings{
			FileExtension:              "pdf",
			FileBaseName:               "",
			IncludeFullPageReceiptsPdf: false,
		},
		Test:         "true",
		OnFinish:     nil,
		isConfigured: false,
	}

	err := myReq.validate()
	assert.Nil(t, err)

	jBytes, jString, err := myReq.marhsall()

	assert.Nil(t, err)
	assert.NotNil(t, jBytes)
	assert.NotEmpty(t, jString)

	expectedJSONString := "{\"type\":\"file\",\"credentials\":{\"partnerUserID\":\"myId\",\"partnerUserSecret\":\"myPassword\"},\"inputSettings\":{\"type\":\"\",\"filters\":{\"reportIDList\":\"\",\"policyIDList\":\"\",\"startDate\":\"2023-01-01\",\"endDate\":\"\",\"approvedAfter\":\"\",\"markedAsExported\":\"\"},\"reportState\":\"\",\"Limit\":\"100\",\"employeeEmail\":\"\"},\"outputSettings\":{\"fileExtension\":\"pdf\",\"fileBaseName\":\"\",\"includeFullPageReceiptsPdf\":false},\"test\":\"true\",\"OnFinish\":null}"

	assert.Equal(t, expectedJSONString, jString)

}
