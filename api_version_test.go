package redmine_test

import (
	"os"
	"testing"
)

func TestGetVersion(t *testing.T) {
	items, err := client.GetVersions(os.Getenv("REDMINE_TEST_PROJECT"), nil, 3)
	if err != nil {
		t.Error(err)
	}

	for _, item := range items {
		if i, err := client.GetVersionByID(item.ID); err != nil {
			t.Error(err)
		} else {
			t.Log(i.ID, i.Name)
		}
	}
}
