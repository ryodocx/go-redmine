package redmine

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"
)

var (
	listLimit int
)

type Option func(*Options)
type Options struct {
	debugMode  bool
	httpClient *http.Client
	listLimit  int
	switchUser string
}

func OptionDebugMode(enabled bool) Option {
	return func(args *Options) {
		args.debugMode = enabled
	}
}
func OptionHTTPClient(client *http.Client) Option {
	return func(args *Options) {
		args.httpClient = client
	}
}
func OptionListLimit(limit int) Option {
	return func(args *Options) {
		args.listLimit = limit
	}
}
func OptionSwitchUser(user string) Option {
	return func(args *Options) {
		args.switchUser = user
	}
}

type Client struct {
	baseURL url.URL
	apiKey  string
	options *Options
}

func NewClient(baseURL, apiKey string, opts ...Option) (*Client, error) {

	o := &Options{
		debugMode:  false,
		httpClient: &http.Client{},
		listLimit:  100,
		switchUser: "",
	}
	for _, setter := range opts {
		setter(o)
	}

	listLimit = o.listLimit

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid baseURL: %s", err.Error())
	}

	c := &Client{
		baseURL: *u,
		apiKey:  apiKey,
		options: o,
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
	if u := c.options.switchUser; u != "" {
		req.Header.Set("X-Redmine-Switch-User", u)
	}

	if c.options.debugMode {
		fmt.Print(req.URL.String())
	}

	resp, err := c.options.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if c.options.debugMode {
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
	_, err := c.getRequest("/my/account.json", url.Values{})
	return err
}
