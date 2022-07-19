package redmine

import (
	"encoding/json"
	"fmt"
)

type Version struct {
	ID             int           `json:"id"`
	Project        IdName        `json:"project"`
	Name           string        `json:"name"`
	Description    string        `json:"description"`
	Status         string        `json:"status"`
	DueDate        *string       `json:"due_date"`
	Sharing        string        `json:"sharing"`
	WikiPageTitle  string        `json:"wiki_page_title"`
	EstimatedHours *float64      `json:"estimated_hours"`
	SpentHours     float64       `json:"spent_hours"`
	CustomFields   []CustomField `json:"custom_fields"`
	CreatedOn      string        `json:"created_on"`
	UpdatedOn      string        `json:"updated_on"`
}

func (c *Client) GetVersionByID(id int, opts ...ReqOption) (*Version, error) {
	endpoint := fmt.Sprintf("/versions/%d.json", id)

	o := reqOptions(opts...)
	respBodyBytes, err := c.getRequest(endpoint, o.query)
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

func (c *Client) GetProjectVersions(project string, opts ...ReqOption) ([]*Version, error) {
	endpoint := fmt.Sprintf("/projects/%s/versions.json", project)

	o := reqOptions(opts...)

	items := []*Version{}
	filter := &listFilter{query: o.query}

	respBodyBytes, err := c.getRequest(endpoint, filter.encode())
	if err != nil {
		return nil, err
	}
	resp := struct {
		listResponseAttrs
		Versions []*Version `json:"versions"`
	}{}
	if err := json.Unmarshal(respBodyBytes, &resp); err != nil {
		return nil, fmt.Errorf("json unmarshal error: %v", err)
	}

	for _, item := range resp.Versions {
		items = append(items, item)
		if o.limit > 0 && len(items) >= o.limit {
			goto end
		}
	}

end:
	return items, nil
}
