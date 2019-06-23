package redmine

import (
	"encoding/json"
	"fmt"
	"time"
)

type Version struct {
	ID           int           `json:"id"`
	Project      IdName        `json:"project"`
	Name         string        `json:"name"`
	Description  string        `json:"description"`
	Status       string        `json:"status"`
	DueDate      string        `json:"due_date"`
	Sharing      string        `json:"sharing"`
	CreatedOn    time.Time     `json:"created_on"`
	UpdatedOn    time.Time     `json:"updated_on"`
	CustomFields []CustomField `json:"custom_fields"`
}

func (c *Client) GetVersion(id int) (*Version, error) {
	endpoint := fmt.Sprintf("/versions/%d.json", id)
	respBodyBytes, err := c.getRequest(endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := struct {
		Version Version `json:"version"`
	}{}
	if err := json.Unmarshal(respBodyBytes, &resp); err != nil {
		return nil, fmt.Errorf("json unmarshal error: %v", err)
	}

	return &resp.Version, nil
}
