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

// callExpensifyExportEndPoint calls any endpoint of Expensify. 'paramsJsonFormat' is what they call 'requestJobDescription' in the API docs
func callExpensifyExportEndPoint(ctx context.Context, requestParams []byte, reportFields []string, expenseFields []string) ([]byte, error) {
	template := generateFreeMarkerTemplate(reportFields, expenseFields)

	// Prepare URL-encoded form data
	formData := url.Values{}
	formData.Add("requestJobDescription", string(requestParams))
	formData.Add("template", template)

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
