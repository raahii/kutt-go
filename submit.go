package kutt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type SubmitInput struct {
	URL       string
	CustomURL string
	Password  string
	Reuse     bool
}

func (cli *Client) Submit(s *SubmitInput) (*URL, error) {
	path := "/api/url/submit"
	reqURL := cli.BaseURL + path

	payload := struct {
		URL       *string `json:"target"`
		CustomURL *string `json:"customurl,omitempty"`
		Password  *string `json:"password,omitempty"`
		Reuse     *bool   `json:"reuse,omitempty"`
	}{
		URL: &s.URL,
	}

	if s.CustomURL != "" {
		payload.CustomURL = &s.CustomURL
	}

	if s.Password != "" {
		payload.Password = &s.Password
	}

	if s.Reuse {
		payload.Reuse = &s.Reuse
	}

	jsonBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshal json: %w", err)
	}

	body := strings.NewReader(string(jsonBytes))
	req, err := http.NewRequest(http.MethodPost, reqURL, body)
	if err != nil {
		return nil, fmt.Errorf("create HTTP request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := cli.do(req)
	if err != nil {
		return nil, fmt.Errorf("do HTTP request: %w", err)
	}

	defer resp.Body.Close()

	if !(resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices) {
		return nil, fmt.Errorf("HTTP response: %w", cli.error(resp.StatusCode, resp.Body))
	}

	var u URL
	if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
		return nil, fmt.Errorf("parse HTTP body: %w", err)
	}

	return &u, nil
}
