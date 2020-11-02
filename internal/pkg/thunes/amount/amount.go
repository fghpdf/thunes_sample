package amount

import "encoding/json"

type BaseModel struct {
	Amount   json.Number `json:"amount,omitempty"` // amount
	Currency string      `json:"currency"`         // reference ISO 4217
}

type SourceModel struct {
	CountryIsoCode string `json:"country_iso_code"` // reference ISO 3166-1 alpha-3
	BaseModel
}

type DestinationModel struct {
	BaseModel
}

type SentModel struct {
	BaseModel
}

type FeeModel struct {
	BaseModel
}
