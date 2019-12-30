package kutt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type DeleteInput struct {
	ID     string
	Domain string
}

func (cli *Client) Delete(s *DeleteInput) error {
	path := "/api/url/deleteurl"
	reqURL := cli.BaseURL + path

	payload := struct {
		ID     *string `json:"id"`
		Domain *string `json:"domain,omitempty"`
	}{
		ID: &s.ID,
	}

	if s.Domain != "" {
		payload.Domain = &s.Domain
	}

	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("marshal json: %w", err)
	}

	body := strings.NewReader(string(jsonBytes))
	req, err := http.NewRequest(http.MethodPost, reqURL, body)
	if err != nil {
		return fmt.Errorf("create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := cli.do(req)
	if err != nil {
		return fmt.Errorf("do HTTP request: %w", err)
	}

	defer resp.Body.Close()

	if !(resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices) {
		return fmt.Errorf("HTTP response: %w", cli.error(resp.StatusCode, resp.Body))
	}

	return nil
}
