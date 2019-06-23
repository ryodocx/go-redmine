package redmine

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

var (
	DebugMode  = false
	HTTPClient = &http.Client{}
	ListLimit  = 100
)

type Client struct {
	baseURL    url.URL
	apiKey     string
	httpClient *http.Client
}

func NewClient(baseURL, apiKey string) (*Client, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid baseURL: %s", err.Error())
	}

	c := &Client{
		baseURL:    *u,
		apiKey:     apiKey,
		httpClient: HTTPClient,
	}

	if err := c.HealthCheck(); err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) getRequest(endpoint string, query url.Values) ([]byte, error) {
	u := c.baseURL
	u.Path = path.Join(u.Path, endpoint)
	if query != nil {
		u.RawQuery = query.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Redmine-API-Key", c.apiKey)

	if DebugMode {
		fmt.Print(req.URL.String())
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if DebugMode {
		fmt.Print(" --> ", resp.Status, "\n")
	}

	switch resp.StatusCode {
	case http.StatusOK:
		// OK
	default:
		return nil, fmt.Errorf("request failed: StatusCode=%d", resp.StatusCode)
	}

	return body, nil
}

func (c *Client) HealthCheck() error {
	query := url.Values{}
	query.Set("limit", "1")
	_, err := c.getRequest("/projects.json", query)
	return err
}
