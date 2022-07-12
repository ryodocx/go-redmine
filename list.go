package redmine

import (
	"net/url"
	"strconv"
)

type listResponseAttrs struct {
	TotalCount int `json:"total_count"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
}

type listFilter struct {
	offset int
	query  url.Values
}

func (f *listFilter) encode() url.Values {
	u := url.Values{}

	if f.query != nil {
		u = f.query
	}

	u.Set("limit", strconv.Itoa(ListLimit))
	u.Set("offset", strconv.Itoa(f.offset))

	return u
}

func (f *listFilter) nextPage() {
	f.offset += ListLimit
}
