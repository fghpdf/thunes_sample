package creditParty

type InfoModel struct {
	CreditPartyIdentifiersAccepted [][]string `json:"credit_party_identifiers_accepted"`
	RequiredReceivingEntityFields  [][]string `json:"required_receiving_entity_fields"`
}
