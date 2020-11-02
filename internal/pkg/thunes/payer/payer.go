package payer

import (
	"encoding/json"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/common"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/service"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/transaction"
)

type Model struct {
	Id               int                   `json:"id"`
	Name             string                `json:"name"`
	Precision        int                   `json:"precision"` // Number of digits after decimal point
	Increment        json.Number           `json:"increment"` // Unit of increment for transaction amounts
	Currency         string                `json:"currency"`  // the currency payer can pay
	CountryIsoCode   string                `json:"country_iso_code"`
	Service          service.Model         `json:"service"`
	TransactionTypes transaction.TypeModel `json:"transaction_types"` // List of transaction types supported with relevant information
}

type ListParams struct {
	common.PageParams
	ServiceId      int    `json:"service_id,omitempty" url:"service_id,omitempty"`             // the ID of available service
	CountryIsoCode string `json:"country_iso_code,omitempty" url:"country_iso_code,omitempty"` // reference ISO 3166-1 alpha-3
	Currency       string `json:"currency,omitempty" url:"currency,omitempty"`                 // reference ISO 4217
}
