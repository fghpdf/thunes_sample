package country

import "fghpdf.me/thunes_homework/internal/pkg/thunes/common"

type Model struct {
	IsoCode string `json:"iso_code"`
	Name    string `json:"name"`
}

type ListParams struct {
	common.PageParams
}
