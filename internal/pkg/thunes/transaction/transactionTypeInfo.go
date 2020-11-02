package transaction

import (
	"encoding/json"
	"fghpdf.me/thunes_homework/internal/pkg/thunes/creditParty"
)

type InfoModel struct {
	MinimumTransactionAmount       json.Number                   `json:"minimum_transaction_amount,omitempty"`
	MaximumTransactionAmount       json.Number                   `json:"maximum_transaction_amount,omitempty"`
	CreditPartyIdentifiersAccepted [][]string                    `json:"credit_party_identifiers_accepted"`
	RequiredSendingEntityFields    [][]string                    `json:"required_sending_entity_fields"`
	RequiredReceivingEntityFields  [][]string                    `json:"required_receiving_entity_fields"`
	RequiredDocuments              [][]string                    `json:"required_documents"`
	CreditPartyInformation         creditParty.InfoModel         `json:"credit_party_information"`
	CreditPartyVerification        creditParty.VerificationModel `json:"credit_party_verification"`
}
