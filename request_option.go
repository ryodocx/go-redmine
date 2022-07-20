package redmine

import "net/url"

type ReqOptions struct {
	query url.Values
	limit int
}
type ReqOption func(*ReqOptions)

func ReqOptionQuery(query url.Values) ReqOption {
	return func(arg *ReqOptions) {
		arg.query = query
	}
}
func ReqOptionLimit(limit int) ReqOption {
	return func(arg *ReqOptions) {
		arg.limit = limit
	}
}

func reqOptions(opts ...ReqOption) *ReqOptions {
	o := &ReqOptions{
		query: url.Values{},
		limit: 0,
	}
	for _, setter := range opts {
		setter(o)
	}
	return o
}
