package transactionType

import (
	"encoding/json"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/creditParty"
)

type InfoModel struct {
	MinimumTransactionAmount       json.Number                    `json:"minimum_transaction_amount,omitempty"`
	MaximumTransactionAmount       json.Number                    `json:"maximum_transaction_amount,omitempty"`
	CreditPartyIdentifiersAccepted [][]string                     `json:"credit_party_identifiers_accepted,omitempty"`
	RequiredSendingEntityFields    [][]string                     `json:"required_sending_entity_fields,omitempty"`
	RequiredReceivingEntityFields  [][]string                     `json:"required_receiving_entity_fields,omitempty"`
	RequiredDocuments              [][]string                     `json:"required_documents,omitempty"`
	CreditPartyInformation         *creditParty.InfoModel         `json:"credit_party_information,omitempty"`
	CreditPartyVerification        *creditParty.VerificationModel `json:"credit_party_verification,omitempty"`
}
