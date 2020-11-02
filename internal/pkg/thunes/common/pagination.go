package common

type PageParams struct {
	Page    int `json:"page,omitempty" url:"page,omitempty"`         // page number
	PerPage int `json:"per_page,omitempty" url:"per_page,omitempty"` // number of results per page, default 50, max 100
}
