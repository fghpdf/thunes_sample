package service

import "fghpdf.me/thunes_homework/internal/pkg/thunes/common"

type Model struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ListParams struct {
	common.PageParams
	CountryIsoCode string `json:"country_iso_code,omitempty" json:"country_iso_code,omitempty"` // reference ISO 3166-1 alpha-3
}
