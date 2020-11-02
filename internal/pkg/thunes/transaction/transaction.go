package transaction

import (
	"encoding/json"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/amount"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/creditParty"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/payer"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/user"
)

type Model struct {
	Id                        uint64                      `json:"id"`
	Status                    string                      `json:"status"`
	StatusMessage             string                      `json:"status_message"`
	StatusClass               string                      `json:"status_class"`
	StatusClassMessage        string                      `json:"status_class_message"`
	ExternalId                string                      `json:"external_id"`
	ExternalCode              string                      `json:"external_code"`
	TransactionType           string                      `json:"transaction_type"`
	PayerTransactionReference string                      `json:"payer_transaction_reference"`
	PayerTransactionCode      string                      `json:"payer_transaction_code"`
	CreationDate              string                      `json:"creation_date"`   // Creation date in HTTP format
	ExpirationDate            string                      `json:"expiration_date"` // Expiration date in HTTP format
	CreditPartyIdentifier     creditParty.IdentifierModel `json:"credit_party_identifier"`
	Source                    amount.SourceModel          `json:"source"`
	Destination               amount.DestinationModel     `json:"destination"`
	Payer                     payer.Model                 `json:"payer"`
	Sender                    user.SenderModel            `json:"sender"`
	Beneficiary               user.BeneficiaryModel       `json:"beneficiary"`
	SendingBusiness           user.BusinessModel          `json:"sending_business"`
	ReceivingBusiness         user.BusinessModel          `json:"receiving_business"`
	CallbackUrl               string                      `json:"callback_url"`
	SentAmount                amount.SentModel            `json:"sent_amount"`
	WholesaleFxRate           json.Number                 `json:"wholesale_fx_rate"`
	RetailRate                json.Number                 `json:"retail_rate"`
	RetailFee                 json.Number                 `json:"retail_fee"`
	RetailFeeCurrency         string                      `json:"retail_fee_currency"` // Retail fee currency in ISO 4217 format
	Fee                       amount.FeeModel             `json:"fee"`
	PurposeOfRemittance       string                      `json:"purpose_of_remittance"`
	DocumentReferenceNumber   string                      `json:"document_reference_number"` // Reference number of document related to the transaction
	AdditionalInformation1    string                      `json:"additional_information_1"`
	AdditionalInformation2    string                      `json:"additional_information_2"`
	AdditionalInformation3    string                      `json:"additional_information_3"`
}

type CreateParams struct {
	CreditPartyIdentifier   creditParty.IdentifierModel `json:"credit_party_identifier"` // required, Credit party identifier
	RetailRate              json.Number                 `json:"retail_rate,omitempty"`
	RetailFee               json.Number                 `json:"retail_fee,omitempty"`
	RetailFeeCurrency       string                      `json:"retail_fee_currency,omitempty"` // Retail fee currency in ISO 4217 format
	Sender                  user.SenderModel            `json:"sender,omitempty"`              // C2C C2B required, Sender information
	Beneficiary             user.BeneficiaryModel       `json:"beneficiary,omitempty"`         // C2C B2C required, Beneficiary information
	SendingBusiness         user.BusinessModel          `json:"sending_business,omitempty"`    // B2C B2B required, Sending business information
	ReceivingBusiness       user.BusinessModel          `json:"receiving_business,omitempty"`  // C2B B2B required, Receiving business information
	ExternalId              string                      `json:"external_id"`                   // required, External ID
	ExternalCode            string                      `json:"external_code,omitempty"`
	CallbackUrl             string                      `json:"callback_url,omitempty"`
	PurposeOfRemittance     string                      `json:"purpose_of_remittance,omitempty"`     // B2B required, Purpose of the remittance
	DocumentReferenceNumber string                      `json:"document_reference_number,omitempty"` // B2B required, Reference number of document related to the transaction
	AdditionalInformation1  string                      `json:"additional_information_1,omitempty"`
	AdditionalInformation2  string                      `json:"additional_information_2,omitempty"`
	AdditionalInformation3  string                      `json:"additional_information_3,omitempty"`
}
