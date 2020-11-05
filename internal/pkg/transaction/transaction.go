package transaction

import (
	"fghpdf.me/thunes_homework/internal/pkg/thunes/creditParty"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/user"
)

type ViewModel struct {
	Id                 uint64 `json:"id"`
	Status             string `json:"status"`
	StatusMessage      string `json:"status_message"`
	StatusClass        string `json:"status_class"`
	StatusClassMessage string `json:"status_class_message"`
	ExternalId         string `json:"external_id"`
	ExternalCode       string `json:"external_code"`
	TransactionType    string `json:"transaction_type"`
	CreationDate       string `json:"creation_date"`   // Creation date in HTTP format
	ExpirationDate     string `json:"expiration_date"` // Expiration date in HTTP format
}

type ViewCreateParams struct {
	QuotationId           uint64                      `json:"quotation_id" uri:"quotationId"`
	CreditPartyIdentifier creditParty.IdentifierModel `json:"credit_party_identifier" form:"credit_party_identifier"` // required, Credit party identifier
	Sender                user.SenderModel            `json:"sender" form:"sender"`                                   // C2C C2B required, Sender information
	Beneficiary           user.BeneficiaryModel       `json:"beneficiary" form:"beneficiary"`                         // C2C B2C required, Beneficiary information
	ExternalId            string                      `json:"external_id" form:"external_id"`                         // required, External ID
}
