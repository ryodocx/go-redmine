package redmine

import (
	"encoding/json"
	"fmt"
	"time"
)

type Project struct {
	ID           int           `json:"id"`
	Name         string        `json:"name"`
	Identifier   string        `json:"identifier"`
	Description  string        `json:"description"`
	Homepage     string        `json:"homepage"`
	Parent       IdName        `json:"parent"`
	Status       int           `json:"status"`
	IsPublic     bool          `json:"is_public"`
	CreatedOn    time.Time     `json:"created_on"`
	UpdatedOn    time.Time     `json:"updated_on"`
	CustomFields []CustomField `json:"custom_fields"`
}

func (c *Client) GetProject(id int) (*Project, error) {
	endpoint := fmt.Sprintf("/projects/%d.json", id)
	respBodyBytes, err := c.getRequest(endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := struct {
		Project Project `json:"project"`
	}{}
	if err := json.Unmarshal(respBodyBytes, &resp); err != nil {
		return nil, fmt.Errorf("json unmarshal error: %v", err)
	}

	return &resp.Project, nil
}

func (c *Client) GetAllProjects() (map[int]*Project, error) {
	endpoint := fmt.Sprintf("/projects.json")

	items := map[int]*Project{}
	resp := struct {
		listResponseAttrs
		Projects []*Project `json:"projects"`
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
		if len(resp.Projects) == 0 {
			break
		}
		for _, item := range resp.Projects {
			items[item.ID] = item
		}
	}

	return items, nil
}
