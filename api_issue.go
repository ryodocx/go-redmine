package redmine

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Issue struct {
	ID                  int           `json:"id"`
	Project             IdName        `json:"project"`
	Tracker             IdName        `json:"tracker"`
	Status              IdName        `json:"status"`
	Priority            IdName        `json:"priority"`
	Author              IdName        `json:"author"`
	AssignedTo          *IdName       `json:"assigned_to,omitempty"`
	FixedVersion        *IdName       `json:"fixed_version,omitempty"`
	Parent              *IdName       `json:"parent,omitempty"`
	Subject             string        `json:"subject"`
	Description         string        `json:"description"`
	StartDate           *string       `json:"start_date,omitempty"`
	DueDate             *string       `json:"due_date,omitempty"`
	DoneRatio           int           `json:"done_ratio"`
	IsPrivate           bool          `json:"is_private"`
	EstimatedHours      *float64      `json:"estimated_hours,omitempty"`
	TotalEstimatedHours *float64      `json:"total_estimated_hours,omitempty"`
	SpentHours          float64       `json:"spent_hours"`
	TotalSpentHours     float64       `json:"total_spent_hours"`
	CustomFields        []CustomField `json:"custom_fields"`
	CreatedOn           string        `json:"created_on"`
	UpdatedOn           string        `json:"updated_on"`
	ClosedOn            *string       `json:"closed_on,omitempty"`
}

func (c *Client) GetIssueByID(id int) (*Issue, error) {
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

func (c *Client) GetIssues(query url.Values, limit int) ([]*Issue, error) {
	endpoint := "/issues.json"

	items := []*Issue{}
	resp := struct {
		listResponseAttrs
		Issues []*Issue `json:"issues"`
	}{}

	filter := &listFilter{query: query}
	for ; ; filter.nextPage() {
		respBodyBytes, err := c.getRequest(endpoint, filter.encode())
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
			items = append(items, item)
			if limit > 0 && len(items) >= limit {
				goto end
			}
		}
	}

end:

	return items, nil
}
