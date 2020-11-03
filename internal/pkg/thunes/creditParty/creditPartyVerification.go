package creditParty

type VerificationModel struct {
	CreditPartyIdentifiersAccepted [][]string `json:"credit_party_identifiers_accepted,omitempty"`
	RequiredReceivingEntityFields  [][]string `json:"required_receiving_entity_fields,omitempty"`
}
