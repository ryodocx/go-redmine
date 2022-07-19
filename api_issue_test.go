package redmine_test

import (
	"testing"

	"github.com/ryodocx/go-redmine/v2"
)

func TestGetIssue(t *testing.T) {
	items, err := client.GetIssues(redmine.ReqOptionLimit(3))
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
