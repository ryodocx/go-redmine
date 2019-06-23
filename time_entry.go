package redmine

import (
	"encoding/json"
	"fmt"
	"time"
)

type TimeEntry struct {
	ID           int           `json:"id"`
	Project      IdName        `json:"project"`
	Issue        Id            `json:"issue"`
	User         IdName        `json:"user"`
	Activity     IdName        `json:"activity"`
	Hours        float64       `json:"hours"`
	Comments     string        `json:"comments"`
	SpentOn      string        `json:"spent_on"`
	CreatedOn    time.Time     `json:"created_on"`
	UpdatedOn    time.Time     `json:"updated_on"`
	CustomFields []CustomField `json:"custom_fields"`
}

func (c *Client) GetTimeEntry(id int) (*TimeEntry, error) {
	endpoint := fmt.Sprintf("/time_entries/%d.json", id)
	respBodyBytes, err := c.getRequest(endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := struct {
		TimeEntry TimeEntry `json:"time_entry"`
	}{}
	if err := json.Unmarshal(respBodyBytes, &resp); err != nil {
		return nil, fmt.Errorf("json unmarshal error: %v", err)
	}

	return &resp.TimeEntry, nil
}

func (c *Client) GetAllTimeEntries() (map[int]*TimeEntry, error) {
	endpoint := fmt.Sprintf("/time_entries.json")

	items := map[int]*TimeEntry{}
	resp := struct {
		listResponseAttrs
		TimeEntries []*TimeEntry `json:"time_entries"`
	}{}

	filter := &listFilter{}
	for ; ; filter.nextPage() {
		respBodyBytes, err := c.getRequest(endpoint, filter.query())
		if err != nil {
			return nil, err
		}
		if err := json.Unmarshal(respBodyBytes, &resp); err != nil {
			return nil, fmt.Errorf("json unmarshal error: %v", err)
		}
		if len(resp.TimeEntries) == 0 {
			break
		}
		for _, item := range resp.TimeEntries {
			items[item.ID] = item
		}
	}

	return items, nil
}
