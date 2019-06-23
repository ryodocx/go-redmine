package redmine_test

import (
	"testing"
)

func TestGetUser(t *testing.T) {
	items, err := client.GetAllUsers()
	if err != nil {
		t.Error(err)
	}

	for k, _ := range items {
		if _, err := client.GetUser(k); err != nil {
			t.Error(err)
		}
		break
	}
}
