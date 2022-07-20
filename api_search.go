package redmine

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Result struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Datetime    string `json:"datetime"`
}

func (c *Client) Search(query []string, opts ...ReqOption) ([]*Result, error) {
	endpoint := "/search.json"

	if query == nil {
		return nil, fmt.Errorf("query is nil")
	}

	o := reqOptions(opts...)
	o.query.Set("q", strings.Join(query, " "))

	items := []*Result{}
	filter := &listFilter{query: o.query}
	for ; ; filter.nextPage() {
		respBodyBytes, err := c.getRequest(endpoint, filter.encode())
		if err != nil {
			return nil, err
		}
		resp := struct {
			listResponseAttrs
			Results []*Result `json:"results"`
		}{}
		if err := json.Unmarshal(respBodyBytes, &resp); err != nil {
			return nil, fmt.Errorf("json unmarshal error: %v", err)
		}
		if len(resp.Results) == 0 {
			goto end
		}
		for _, item := range resp.Results {
			items = append(items, item)
			if o.limit > 0 && len(items) >= o.limit {
				goto end
			}
		}
	}

end:
	return items, nil
}
