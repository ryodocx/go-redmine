package redmine

import (
	"encoding/json"
	"fmt"
	"time"
)

type User struct {
	ID          int       `json:"id"`
	Login       string    `json:"login"`
	Firstname   string    `json:"firstname"`
	Lastname    string    `json:"lastname"`
	Mail        string    `json:"mail"`
	CreatedOn   time.Time `json:"created_on"`
	LastLoginOn time.Time `json:"last_login_on"`
	// APIKey      string    `json:"api_key"`
	Status       int           `json:"status"`
	CustomFields []CustomField `json:"custom_fields"`
}

func (c *Client) GetUser(id int) (*User, error) {
	endpoint := fmt.Sprintf("/users/%d.json", id)
	respBodyBytes, err := c.getRequest(endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := struct {
		User User `json:"user"`
	}{}
	if err := json.Unmarshal(respBodyBytes, &resp); err != nil {
		return nil, fmt.Errorf("json unmarshal error: %v", err)
	}

	return &resp.User, nil
}

func (c *Client) GetAllUsers() (map[int]*User, error) {
	endpoint := fmt.Sprintf("/users.json")

	items := map[int]*User{}
	resp := struct {
		listResponseAttrs
		Users []*User `json:"users"`
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
		if len(resp.Users) == 0 {
			break
		}
		for _, item := range resp.Users {
			items[item.ID] = item
		}
	}

	return items, nil
}
