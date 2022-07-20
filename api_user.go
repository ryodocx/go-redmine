package redmine

type User struct {
	ID              int           `json:"id"`
	Login           string        `json:"login"`
	Admin           bool          `json:"admin"`
	Firstname       string        `json:"firstname"`
	Lastname        string        `json:"lastname"`
	Mail            string        `json:"mail"`
	CreatedOn       string        `json:"created_on"`
	LastLoginOn     string        `json:"last_login_on"`
	PasswdChangedOn *string       `json:"passwd_changed_on"`
	TwofaScheme     *string       `json:"twofa_scheme"`
	APIKey          string        `json:"api_key"`
	CustomFields    []CustomField `json:"custom_fields"`
}
