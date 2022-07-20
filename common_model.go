package redmine

type Id struct {
	Id int `json:"id"`
}

type IdName struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CustomField struct {
	ID    int         `json:"id"`
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}
