package kutt

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type StatsInput struct {
	ID     string
	Domain string
}

func (cli *Client) Stats(s *StatsInput) error {
	path := "/api/url/stats"

	params := url.Values{}
	params.Add("id", s.ID)

	if s.Domain != "" {
		params.Add("domain", s.Domain)
	}

	reqURL := cli.BaseURL + path + "?" + params.Encode()
	fmt.Println(reqURL)
	req, err := http.NewRequest(http.MethodGet, reqURL, nil)
	if err != nil {
		return fmt.Errorf("create HTTP request: %w", err)
	}

	resp, err := cli.do(req)
	if err != nil {
		return fmt.Errorf("do HTTP request: %w", err)
	}

	defer resp.Body.Close()

	if !(resp.StatusCode >= http.StatusOK && resp.StatusCode < http.StatusMultipleChoices) {
		return fmt.Errorf("HTTP response: %w", cli.error(resp.StatusCode, resp.Body))
	}

	var respBody io.Reader = resp.Body

	// For Debug
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(respBody); err != nil {
		return fmt.Errorf("read HTTP body: %w", err)
	}
	fmt.Println(buf.String())

	return nil
}
