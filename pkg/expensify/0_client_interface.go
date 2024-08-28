package expensify

import (
	"context"
	"encoding/json"
	"fmt"
)

// Client is a struct to hold the ExpensifyClient information
type Client struct {
	partnerUserID     string
	partnerUserSecret string
}

// NewClient creates a new ExpensifyClient
func NewClient(partnerUserID, partnerUserSecret string) (*Client, error) {
	return &Client{
		partnerUserID:     partnerUserID,
		partnerUserSecret: partnerUserSecret,
	}, nil
}

func (c *Client) GetReportsInJson(ctx context.Context,
	config FileExportBaseConfig,
	sendEmailOnFinish *OnFinishSendEmailConfig,
	markAsExportedOnFinish *OnFinishMarkAsExportedConfig,
	uploadToSftpOnFinish *OnFinishSftpUploadDataConfig,
	reportFieldNames []string,
	expenseFieldNames []string) ([]Report, error) {

	// ----------- START

	//  configure the export
	config.outputFileExtension = "json"
	req := fileExportRequest{}
	err := req.init(config, sendEmailOnFinish, markAsExportedOnFinish, uploadToSftpOnFinish, c.partnerUserID, c.partnerUserSecret)
	if err != nil {
		return nil, fmt.Errorf("error while initializing fileExportRequest: %w", err)
	}

	paramsJsonFormat, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling the exportrequest: %w", err)
	}

	// execute the export - now the file is created on the server and we get back the file name
	inBytes, err := callExpensifyExportEndPoint(ctx, paramsJsonFormat, reportFieldNames, expenseFieldNames)
	if err != nil {
		return nil, fmt.Errorf("error while calling the ExpensifyEndpoint: %w", err)
	}

	// configuring the download - we download the file from the server
	dwnConfig := downloaderConfig{
		Type: "download",
		Credentials: expCredentials{
			PartnerUserID:     c.partnerUserID,
			PartnerUserSecret: c.partnerUserSecret,
		},
		FileName:   string(inBytes),
		FileSystem: "integrationServer",
	}
	configJson, err := json.Marshal(dwnConfig)
	if err != nil {
		return nil, fmt.Errorf("error while marshalling the exportrequest: %w", err)
	}

	inBytes, err = callExpensifyExportEndPoint(ctx, configJson, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error while calling expenisfyEdnpoint: %w", err)
	}

	// Turning the [] of JSON objects into individual JSON objects per report

	// Step 1: Validate if inBytes is a valid JSON array
	var jsonArray []json.RawMessage
	err = json.Unmarshal(inBytes, &jsonArray)
	if err != nil {
		return nil, fmt.Errorf("received data is not a valid JSON array: %w", err)
	}

	if len(jsonArray) == 0 {
		return nil, fmt.Errorf("received JSON array is empty")
	}

	// Step 2: Convert JSON array to JSON object by removing the square brackets
	outReports := make([]Report, 0)
	for _, jsonData := range jsonArray {
		myReport := Report{
			FileExtension: "json",
			Data:          jsonData,
		}
		outReports = append(outReports, myReport)

		// Step 3: Validate if the resulting byte slice is a valid JSON object
		var jsonObject map[string]interface{}
		err = json.Unmarshal(jsonData, &jsonObject)
		if err != nil {
			return nil, fmt.Errorf("modified data is not a valid JSON object: %w", err)
		}

	}

	return outReports, nil
}
