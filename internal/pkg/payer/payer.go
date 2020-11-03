package payer

import "fghpdf.me/thunes_homework/internal/pkg/thunes/transactionType"

type ViewModel struct {
	Id               int                       `json:"id"`
	Name             string                    `json:"name"`
	Currency         string                    `json:"currency"` // the currency payer can pay
	CountryIsoCode   string                    `json:"country_iso_code"`
	TransactionTypes transactionType.TypeModel `json:"transaction_types"` // List of transaction types supported with relevant information
}
