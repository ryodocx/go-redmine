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
	limit  int
	offset int
	others url.Values
}

func (f *listFilter) query() url.Values {
	u := url.Values{}

	if f.limit < 1 || 100 < f.limit {
		f.limit = ListLimit
	}
	u.Set("limit", strconv.Itoa(f.limit))

	u.Set("offset", strconv.Itoa(f.offset))

	if f.others != nil {
		for k, v := range f.others {
			u.Set(k, v[0])
		}
	}

	return u
}

func (f *listFilter) nextPage() {
	f.offset += f.limit
}
