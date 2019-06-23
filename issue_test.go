package redmine_test

import (
	"testing"
)

func TestGetIssue(t *testing.T) {
	items, err := client.GetAllIssues()
	if err != nil {
		t.Error(err)
	}

	for k, _ := range items {
		if _, err := client.GetIssue(k); err != nil {
			t.Error(err)
		}
		break
	}
}
