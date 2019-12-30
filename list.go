package kutt

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type listResponse struct {
	URLs  []*URL `json:"list"`
	Count int    `json:"countAll"`
}

func (cli *Client) List() ([]*URL, error) {
	path := "/api/url/geturls"
	reqURL := cli.BaseURL + path

	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return nil, fmt.Errorf("create HTTP request: %w", err)
	}

	resp, err := cli.do(req)
	if err != nil {
		return nil, fmt.Errorf("do HTTP request: %w", err)
	}

	defer resp.Body.Close()

	if !(resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices) {
		return nil, fmt.Errorf("HTTP response: %w", cli.error(resp.StatusCode, resp.Body))
	}

	var r listResponse
	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		return nil, fmt.Errorf("parse HTTP body: %w", err)
	}

	return r.URLs, nil
}
