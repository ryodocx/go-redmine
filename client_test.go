package redmine_test

import (
	"os"
	"testing"

	"github.com/ryodocx/go-redmine/v2"
)

var (
	client          *redmine.Client
	baseURL, apiKey string
)

func init() {

	redmine.DebugMode = true

	baseURL = os.Getenv("REDMINE_BASEURL")
	if baseURL == "" {
		panic("REDMINE_BASEURL is required!")
	}

	apiKey = os.Getenv("REDMINE_APIKEY")
	if apiKey == "" {
		panic("REDMINE_APIKEY is required!")
	}

	if c, err := redmine.NewClient(baseURL, apiKey); err != nil {
		panic(err)
	} else {
		client = c
	}
}

func Example() {
	client, err := redmine.NewClient(baseURL, apiKey)
	if err != nil {
		panic(err)
	}

	client.HealthCheck()
	// client.GetXxx...
}

func TestHealthCheck(t *testing.T) {
	if err := client.HealthCheck(); err != nil {
		t.Fatal(err)
	}
}
