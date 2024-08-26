package expensify

import (
	"context"
	"encoding/json"
	"fmt"
)

type response struct {
	ResponseMessage string `json:"responseMessage"`
	ResponseCode    int    `json:"responseCode"`
}

// ExecuteFileExport configures the export request. Everything that is marked a point is nullable (meaning it is optional to configure)
// -filterByMarkedAsApprovedTag : pass the tag you are looking for
func (c *Client) ExecuteFileExport(ctx context.Context, reportFields []string, expenseFields []string) ([]*ExpensifyReport, error) {
	myReq := c.fileExportConfig

	jsonData, err := json.Marshal(myReq)
	if err != nil {
		return nil, err
	}
	inBytes, err := callExpensifyEndPoint(context.Background(), jsonData, reportFields, expenseFields)
	if err != nil {
		return nil, fmt.Errorf("error while calling the ExpensifyEndpoint: %w", err)
	}

	// Expensify does not return a []string with the report name, but a simple comma separated string
	// we split this into actual report names
	outReportNames, err := SplitFilenames(string(inBytes))
	if err != nil {
		return nil, fmt.Errorf("error while splitting the names of the incoming reports: %w", err)
	}

	outReports := make([]*ExpensifyReport, 0)

	for _, fileName := range outReportNames {
		outReport := ExpensifyReport{
			FileName:   fileName,
			FileSystem: "integrationServer",
			Data:       nil,
		}
		outReports = append(outReports, &outReport)
	}

	return outReports, nil
}
