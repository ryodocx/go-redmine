package redmine

import (
	"encoding/json"
	"fmt"
	"time"
)

type Issue struct {
	ID                  int           `json:"id"`
	Project             IdName        `json:"project"`
	Tracker             IdName        `json:"tracker"`
	Status              IdName        `json:"status"`
	Priority            IdName        `json:"priority"`
	Author              IdName        `json:"author"`
	FixedVersion        IdName        `json:"fixed_version"`
	Subject             string        `json:"subject"`
	Description         string        `json:"description"`
	StartDate           string        `json:"start_date"`
	DueDate             *string       `json:"due_date"`
	DoneRatio           int           `json:"done_ratio"`
	IsPrivate           bool          `json:"is_private"`
	EstimatedHours      float64       `json:"estimated_hours"`
	TotalEstimatedHours float64       `json:"total_estimated_hours"`
	SpentHours          float64       `json:"spent_hours"`
	TotalSpentHours     float64       `json:"total_spent_hours"`
	CreatedOn           time.Time     `json:"created_on"`
	UpdatedOn           time.Time     `json:"updated_on"`
	ClosedOn            *time.Time    `json:"closed_on"`
	CustomFields        []CustomField `json:"custom_fields"`
}

func (c *Client) GetIssue(id int) (*Issue, error) {
	endpoint := fmt.Sprintf("/issues/%d.json", id)
	respBodyBytes, err := c.getRequest(endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := struct {
		Issue Issue `json:"issue"`
	}{}
	if err := json.Unmarshal(respBodyBytes, &resp); err != nil {
		return nil, fmt.Errorf("json unmarshal error: %v", err)
	}

	return &resp.Issue, nil
}

func (c *Client) GetAllIssues() (map[int]*Issue, error) {
	endpoint := fmt.Sprintf("/issues.json")

	items := map[int]*Issue{}
	resp := struct {
		listResponseAttrs
		Issues []*Issue `json:"issues"`
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
		if len(resp.Issues) == 0 {
			break
		}
		for _, item := range resp.Issues {
			items[item.ID] = item
		}
	}

	return items, nil
}
