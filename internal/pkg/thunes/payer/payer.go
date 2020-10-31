package payer

import (
	"fghpdf.me/thunes_homework/internal/pkg/thunes/service"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/transactionType"
)

type Model struct {
	Id               int                   `json:"id"`
	Name             string                `json:"name"`
	Precision        int                   `json:"precision"` // Number of digits after decimal point
	Increment        string                `json:"increment"` // Unit of increment for transaction amounts
	Currency         string                `json:"currency"`  // the currency payer can pay
	CountryIsoCode   string                `json:"country_iso_code"`
	Service          service.Model         `json:"service"`
	TransactionTypes transactionType.Model `json:"transaction_types"` // List of transaction types supported with relevant information
}

type ListParams struct {
	Page           int    `json:"page"`             // page number
	PerPage        int    `json:"per_page"`         // number of results per page, default 50, max 100
	ServiceId      int    `json:"service_id"`       // the ID of available service
	CountryIsoCode string `json:"country_iso_code"` // reference ISO 3166-1 alpha-3
	Currency       string `json:"currency"`         // reference ISO 4217
}
