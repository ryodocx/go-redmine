package redmine

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetMyAccount() (*User, error) {
	endpoint := "/my/account.json"

	respBodyBytes, err := c.getRequest(endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp := struct {
		MyAccount User `json:"user"`
	}{}
	if err := json.Unmarshal(respBodyBytes, &resp); err != nil {
		return nil, fmt.Errorf("json unmarshal error: %v", err)
	}

	return &resp.MyAccount, nil
}
