package service

type Model struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ListParams struct {
	Page           int    `json:"page,omitempty"`             // page number
	PerPage        int    `json:"per_page,omitempty"`         // number of results per page, default 50, max 100
	CountryIsoCode string `json:"country_iso_code,omitempty"` // reference ISO 3166-1 alpha-3
}
