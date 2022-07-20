package redmine_test

import (
	"testing"
)

func TestGetMyAccount(t *testing.T) {
	_, err := client.GetMyAccount()
	if err != nil {
		t.Error(err)
	}
}
