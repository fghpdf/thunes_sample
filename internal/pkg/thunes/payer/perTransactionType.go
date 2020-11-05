package payer

type PerTransactionTypeModel struct {
	C2C *RateInfoModel `json:"C2C,omitempty"`
	B2B *RateInfoModel `json:"B2B,omitempty"`
	C2B *RateInfoModel `json:"C2B,omitempty"`
	B2C *RateInfoModel `json:"B2C,omitempty"`
}
