package payer

import "encoding/json"

type RateDetailModel struct {
	SourceAmountMin json.Number `json:"source_amount_min"`
	SourceAmountMax json.Number `json:"source_amount_max"`
	WholesaleFxRate json.Number `json:"wholesale_fx_rate"`
}
