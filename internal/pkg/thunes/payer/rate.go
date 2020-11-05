package payer

type RateModel struct {
	DestinationCurrency string        `json:"destination_currency"`
	Rates               RateInfoModel `json:"rates"`
}
