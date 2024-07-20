package expensify

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type response struct {
	ResponseMessage string `json:"responseMessage"`
	ResponseCode    int    `json:"responseCode"`
}

// ExecuteFileExport configures the export request. Everything that is marked a point is nullable (meaning it is optional to configure)
// -filterByMarkedAsApprovedTag : pass the tag you are looking for
func (c *ExpensifyClient) ExecuteFileExport(ctx context.Context) ([]byte, error) {
	myReq := c.fileExportConfig

	jsonData, err := json.Marshal(myReq)
	if err != nil {
		return nil, err
	}

	// Prepare URL-encoded form data
	formData := url.Values{}
	formData.Add("requestJobDescription", string(jsonData))
	formData.Add("template", "expensify_template.ftl")

	myUrl := baseURL + "?" + formData.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, myUrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	clientHTTP := &http.Client{}
	resp, err := clientHTTP.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var myResponse response

	// Expensify has an inconsistent response structure. when there is no error, they do not send back a response code
	if err := json.Unmarshal(body, &myResponse); err != nil {
		return body, nil
	}

	if myResponse.ResponseCode != 200 {
		return nil, fmt.Errorf("%s", body)
	}
	return body, nil
}
