package redmine_test

import (
	"os"
	"testing"

	"github.com/ryodocx/go-redmine"
)

func TestSearch(t *testing.T) {
	items, err := client.Search([]string{os.Getenv("REDMINE_TEST_QUERY")}, redmine.ReqOptionLimit(3))
	if err != nil {
		t.Error(err)
	}

	for _, v := range items {
		t.Log(v.Title)
	}
}
