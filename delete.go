package kutt

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type DeleteParams struct {
	ID     string  `json:"id"`
	Domain *string `json:"domain"`
}

type DeleteOption func(*DeleteParams)

func WithDomain(v string) DeleteOption {
	return func(p *DeleteParams) {
		p.Domain = &v
	}
}

func (cli *Client) Delete(ID string, opts ...DeleteOption) error {
	path := "/api/url/deleteurl"
	reqURL := cli.BaseURL + path

	payload := &DeleteParams{
		ID: ID,
	}

	for _, opt := range opts {
		opt(payload)
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
