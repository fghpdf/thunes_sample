package quotation

import (
	"encoding/json"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/amount"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/payer"
)

type Model struct {
	Id              int                     `json:"id"`
	ExternalId      string                  `json:"external_id"` // External reference ID. This is the reference for this quotation from the external system.
	Payer           payer.Model             `json:"payer"`
	Mode            string                  `json:"mode"` // Quotation Mode
	TransactionType string                  `json:"transaction_type"`
	Source          amount.SourceModel      `json:"source"`
	Destination     amount.DestinationModel `json:"destination"`
	SentAmount      amount.SentModel        `json:"sent_amount"`
	WholesaleFxRate json.Number             `json:"wholesale_fx_rate"` // Wholesale FX rate
	Fee             amount.FeeModel         `json:"fee"`
	CreationDate    string                  `json:"creation_date"`
	ExpirationDate  string                  `json:"expiration_date"`
}

type CreateParams struct {
	ExternalId      string                  `json:"external_id"`      // required, External reference ID
	PayerId         int                     `json:"payer_id"`         // required, Payer ID
	Mode            string                  `json:"mode"`             // required, Quotation Mode
	TransactionType string                  `json:"transaction_type"` // required, transaction type
	Source          amount.SourceModel      `json:"source"`           // required, source information
	Destination     amount.DestinationModel `json:"destination"`      // required, destination information
}
