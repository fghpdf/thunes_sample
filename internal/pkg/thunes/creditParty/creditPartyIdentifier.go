package creditParty

type IdentifierModel struct {
	Msisdn            string `json:"msisdn,omitempty"`
	BankAccountNumber string `json:"bank_account_number,omitempty"`
	Iban              string `json:"iban,omitempty"`
	Clabe             string `json:"clabe,omitempty"`
	Cbu               string `json:"cbu,omitempty"`
	CbuAlias          string `json:"cbu_alias,omitempty"`
	SwiftBicCode      string `json:"swift_bic_code,omitempty"`
	BikCode           string `json:"bik_code,omitempty"`
	IfsCode           string `json:"ifs_code,omitempty"`
	SortCode          string `json:"sort_code,omitempty"`
	AbaRoutingNumber  string `json:"aba_routing_number,omitempty"`
	BsbNumber         string `json:"bsb_number,omitempty"`
	BranchNumber      string `json:"branch_number,omitempty"`
	RoutingCode       string `json:"routing_code,omitempty"`
	EntityTtId        string `json:"entity_tt_id,omitempty"`
	AccountType       string `json:"account_type,omitempty"`
	AccountNumber     string `json:"account_number,omitempty"`
	Email             string `json:"email,omitempty"`
}
