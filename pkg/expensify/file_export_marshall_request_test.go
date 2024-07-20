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

	expectedJSONString := "{\"type\":\"file\",\"onReceive\":{\"immediateResponse\":null},\"credentials\":{\"partnerUserID\":\"myId\",\"partnerUserSecret\":\"myPassword\"},\"inputSettings\":{\"filters\":{\"startDate\":\"2023-01-01\"},\"Limit\":\"100\"},\"outputSettings\":{\"fileExtension\":\"pdf\"},\"test\":\"true\"}"

	assert.Equal(t, expectedJSONString, jString)

}
