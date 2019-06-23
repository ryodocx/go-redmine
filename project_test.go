package redmine_test

import (
	"testing"
)

func TestGetProject(t *testing.T) {
	items, err := client.GetAllProjects()
	if err != nil {
		t.Error(err)
	}

	for k, _ := range items {
		if _, err := client.GetProject(k); err != nil {
			t.Error(err)
		}
		break
	}
}
