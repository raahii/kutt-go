package kutt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type SubmitParams struct {
	URL       string  `json:"target"`
	CustomURL *string `json:"customurl"`
	Password  *string `json:"password"`
	Reuse     *bool   `json:"reuse"`
}

type SubmitOption func(*SubmitParams)

func WithCustomURL(v string) SubmitOption {
	return func(p *SubmitParams) {
		p.CustomURL = &v
	}
}

func WithPassword(v string) SubmitOption {
	return func(p *SubmitParams) {
		p.Password = &v
	}
}

func WithReuse(v bool) SubmitOption {
	return func(p *SubmitParams) {
		p.Reuse = &v
	}
}

func (cli *Client) Submit(target string, opts ...SubmitOption) (*URL, error) {
	path := "/api/url/submit"
	reqURL := cli.BaseURL + path

	payload := &SubmitParams{
		URL: target,
	}

	for _, opt := range opts {
		opt(payload)
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
