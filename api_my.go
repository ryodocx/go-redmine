package redmine

import (
	"encoding/json"
	"fmt"
)

type MyAccount struct {
	ID           int           `json:"id"`
	Login        string        `json:"login"`
	Admin        bool          `json:"admin"`
	APIKey       string        `json:"api_key"`
	Firstname    string        `json:"firstname"`
	Lastname     string        `json:"lastname"`
	CreatedOn    string        `json:"created_on"`
	Mail         string        `json:"mail"`
	LastLoginOn  string        `json:"last_login_on"`
	CustomFields []CustomField `json:"custom_fields"`
}

func (c *Client) GetMyAccount() (*MyAccount, error) {
	endpoint := "/my/account.json"

	respBodyBytes, err := c.getRequest(endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := struct {
		MyAccount MyAccount `json:"user"`
	}{}
	if err := json.Unmarshal(respBodyBytes, &resp); err != nil {
		return nil, fmt.Errorf("json unmarshal error: %v", err)
	}

	return &resp.MyAccount, nil
}
