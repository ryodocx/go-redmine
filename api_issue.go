package redmine

import (
	"encoding/json"
	"fmt"
)

type Issue struct {
	ID                  int           `json:"id"`
	Project             IdName        `json:"project"`
	Tracker             IdName        `json:"tracker"`
	Status              IdName        `json:"status"`
	Priority            IdName        `json:"priority"`
	Author              IdName        `json:"author"`
	AssignedTo          *IdName       `json:"assigned_to"`
	FixedVersion        *IdName       `json:"fixed_version"`
	Parent              *IdName       `json:"parent"`
	Subject             string        `json:"subject"`
	Description         string        `json:"description"`
	StartDate           *string       `json:"start_date"`
	DueDate             *string       `json:"due_date"`
	DoneRatio           int           `json:"done_ratio"`
	IsPrivate           bool          `json:"is_private"`
	EstimatedHours      *float64      `json:"estimated_hours"`
	TotalEstimatedHours *float64      `json:"total_estimated_hours"`
	SpentHours          float64       `json:"spent_hours"`
	TotalSpentHours     float64       `json:"total_spent_hours"`
	CustomFields        []CustomField `json:"custom_fields"`
	CreatedOn           string        `json:"created_on"`
	UpdatedOn           string        `json:"updated_on"`
	ClosedOn            *string       `json:"closed_on"`
}

func (c *Client) GetIssueByID(id int, opts ...ReqOption) (*Issue, error) {
	endpoint := fmt.Sprintf("/issues/%d.json", id)

	o := reqOptions(opts...)
	respBodyBytes, err := c.getRequest(endpoint, o.query)
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

func (c *Client) GetIssues(opts ...ReqOption) ([]*Issue, error) {
	endpoint := "/issues.json"

	o := reqOptions(opts...)

	items := []*Issue{}
	filter := &listFilter{query: o.query}
	for ; ; filter.nextPage() {
		respBodyBytes, err := c.getRequest(endpoint, filter.encode())
		if err != nil {
			return nil, err
		}
		resp := struct {
			listResponseAttrs
			Issues []*Issue `json:"issues"`
		}{}
		if err := json.Unmarshal(respBodyBytes, &resp); err != nil {
			return nil, fmt.Errorf("json unmarshal error: %v", err)
		}
		if len(resp.Issues) == 0 {
			goto end
		}
		for _, item := range resp.Issues {
			items = append(items, item)
			if o.limit > 0 && len(items) >= o.limit {
				goto end
			}
		}
	}

end:
	return items, nil
}
