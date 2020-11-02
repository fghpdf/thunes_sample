package transaction

type TypeModel struct {
	C2C InfoModel `json:"C2C"`
	C2B InfoModel `json:"C2B"`
	B2C InfoModel `json:"B2C"`
	B2B InfoModel `json:"B2B"`
}
