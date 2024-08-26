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
	fileExportConfig  *fileExportRequest
}

// NewClient creates a new ExpensifyClient
func NewClient(partnerUserID, partnerUserSecret string) (*Client, error) {
	return &Client{
		partnerUserID:     partnerUserID,
		partnerUserSecret: partnerUserSecret,
	}, nil
}

// DownloadReport  downloads the data of exported report
func (c *Client) DownloadReport(ctx context.Context, report *ExpensifyReport) error {
	err := report.validate()
	if err != nil {
		return fmt.Errorf("error while validating the ExpensifyReport: %w", err)
	}

	config := downloaderConfig{
		Type: "download",
		Credentials: expCredentials{
			PartnerUserID:     c.partnerUserID,
			PartnerUserSecret: c.partnerUserSecret,
		},
		FileName:   report.FileName,
		FileSystem: report.FileSystem,
	}
	jsonBytes, err := json.Marshal(config)
	if err != nil {
		return fmt.Errorf("error while marshalling downloaderConfig: %w", err)
	}

	inBytes, err := callExpensifyEndPoint(ctx, jsonBytes, nil, nil)
	if err != nil {
		return fmt.Errorf("error while calling expenisfyEdnpoint: %w", err)
	}

	report.Data = inBytes

	return nil
}
