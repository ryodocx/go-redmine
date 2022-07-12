package redmine_test

import (
	"testing"
)

func TestGetIssue(t *testing.T) {
	items, err := client.GetIssues(nil, 3)
	if err != nil {
		t.Error(err)
	}

	for _, item := range items {
		if i, err := client.GetIssueByID(item.ID); err != nil {
			t.Error(err)
		} else {
			t.Log(i.ID, i.Subject)
		}
	}
}
