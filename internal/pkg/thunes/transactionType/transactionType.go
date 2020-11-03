package transactionType

type TypeModel struct {
	C2C *InfoModel `json:"C2C,omitempty"`
	C2B *InfoModel `json:"C2B,omitempty"`
	B2C *InfoModel `json:"B2C,omitempty"`
	B2B *InfoModel `json:"B2B,omitempty"`
}
