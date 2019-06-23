package redmine_test

import (
	"testing"
)

func TestGetTimeEntry(t *testing.T) {
	items, err := client.GetAllTimeEntries()
	if err != nil {
		t.Error(err)
	}

	for k, _ := range items {
		if _, err := client.GetTimeEntry(k); err != nil {
			t.Error(err)
		}
		break
	}
}
