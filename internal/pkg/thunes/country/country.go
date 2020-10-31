package country

type Model struct {
	IsoCode string `json:"iso_code"`
	Name    string `json:"name"`
}

type ListParams struct {
	Page    int `json:"page,omitempty"`     // page number
	PerPage int `json:"per_page,omitempty"` // number of results per page, default 50, max 100
}
