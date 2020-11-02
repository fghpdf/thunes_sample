package user

type BaseUserModel struct {
	Lastname                  string `json:"lastname,omitempty"`
	Lastname2                 string `json:"lastname2,omitempty"`
	MiddleName                string `json:"middlename,omitempty"`
	Firstname                 string `json:"firstname,omitempty"`
	NativeName                string `json:"nativename,omitempty"`
	NationalityCountryIsoCode string `json:"nationality_country_iso_code,omitempty"` // Nationality in ISO 3166-1 alpha-3 format
	Code                      string `json:"code,omitempty"`                         // Sender identification code
	DateOfBirth               string `json:"date_of_birth,omitempty"`                // Date of birth in ISO 8601 format
	CountryOfBirthIsoCode     string `json:"country_of_birth_iso_code,omitempty"`    // Country of birth in ISO 3166-1 alpha-3 format
	Gender                    string `json:"gender,omitempty"`
	Address                   string `json:"address,omitempty"`
	PostalCode                string `json:"postal_code,omitempty"`
	City                      string `json:"city,omitempty"`
	CountryIsoCode            string `json:"country_iso_code,omitempty"` // Address country in ISO 3166-1 alpha-3 format
	Msisdn                    string `json:"msisdn,omitempty"`           // MSISDN in international format
	Email                     string `json:"email,omitempty"`
	IdType                    string `json:"id_type,omitempty"`             // Presented identification type
	IdCountryIsoCode          string `json:"id_country_iso_code,omitempty"` // ID country in ISO 3166-1 alpha-3 format
	IdNumber                  string `json:"id_number,omitempty"`           // Presented identification number
	IdDeliveryDate            string `json:"id_delivery_date,omitempty"`    // ID delivery date in ISO 8601 format
	IdExpirationDate          string `json:"id_expiration_date,omitempty"`  // ID expiration date in ISO 8601 format
	Occupation                string `json:"occupation,omitempty"`
	ProvinceState             string `json:"province_state,omitempty"` // Address province/state
}

type SenderModel struct {
	BaseUserModel
	Bank                    string `json:"bank,omitempty"`                     // Name of bank
	BankAccount             string `json:"bank_account,omitempty"`             // Bank account number
	Card                    string `json:"card,omitempty"`                     // Credit card number
	BeneficiaryRelationship string `json:"beneficiary_relationship,omitempty"` // Relationship to beneficiary
	SourceOfFunds           string `json:"source_of_funds,omitempty"`
}

type BeneficiaryModel struct {
	BaseUserModel
	BankAccountHolderName string `json:"bank_account_holder_name,omitempty"`
}
