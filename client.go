package kutt

import (
	"io"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

const version = "0.0.1"

const baseURL = "https://kutt.it"

func init() {
	defaultUserAgent = "kutt-go/" + version + " (+https://github.com/raahii/kutt-go)"
}

var defaultUserAgent string

type Client struct {
	HTTPClient *http.Client
	ApiKey     string
	BaseURL    string
	UserAgent  string
}

func NewClient(apiKey string) *Client {
	var cli Client
	cli.ApiKey = apiKey
	cli.BaseURL = baseURL
	cli.UserAgent = defaultUserAgent

	return &cli
}

func (cli *Client) getUA() string {
	if cli.UserAgent != "" {
		return cli.UserAgent
	}
	return defaultUserAgent
}

func (cli *Client) httpClient() *http.Client {
	if cli.HTTPClient != nil {
		return cli.HTTPClient
	}
	return http.DefaultClient
}

func (cli *Client) do(req *http.Request) (*http.Response, error) {
	req.Header.Set("X-API-Key", cli.ApiKey)
	req.Header.Set("User-Agent", cli.getUA())
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	return cli.httpClient().Do(req)
}

func (cli *Client) error(statusCode int, body io.Reader) error {
	buf, err := ioutil.ReadAll(body)
	if err != nil || len(buf) == 0 {
		return errors.Errorf("request failed with status code %d", statusCode)
	}
	return errors.Errorf("StatusCode: %d, Error: %s", statusCode, string(buf))
}
