package user

type BusinessModel struct {
	RegisteredName                 string `json:"registered_name,omitempty"`
	TradingName                    string `json:"trading_name,omitempty"`
	Address                        string `json:"address,omitempty"`
	PostalCode                     string `json:"postal_code,omitempty"`
	City                           string `json:"city,omitempty"`
	ProvinceState                  string `json:"province_state,omitempty"`
	CountryIsoCode                 string `json:"country_iso_code,omitempty"` // Country of the business in ISO 3166-1 alpha-3 format
	Msisdn                         string `json:"msisdn,omitempty"`
	Email                          string `json:"email,omitempty"`
	RegistrationNumber             string `json:"registration_number,omitempty"`
	TaxId                          string `json:"tax_id,omitempty"`
	DateOfIncorporation            string `json:"date_of_incorporation,omitempty"`
	RepresentativeLastname         string `json:"representative_lastname,omitempty"`
	RepresentativeLastname2        string `json:"representative_lastname2,omitempty"`
	RepresentativeFirstname        string `json:"representative_firstname,omitempty"`
	RepresentativeMiddleName       string `json:"representative_middlename,omitempty"`
	RepresentativeNativeName       string `json:"representative_nativename,omitempty"`
	RepresentativeIdType           string `json:"representative_id_type,omitempty"`             // Presented identification type of the representative of the business
	RepresentativeIdCountryIsoCode string `json:"representative_id_country_iso_code,omitempty"` // ID delivery date in ISO 8601 format
	RepresentativeIdNumber         string `json:"representative_id_number,omitempty"`           // Presented identification number
	RepresentativeIdDeliveryDate   string `json:"representative_id_delivery_date,omitempty"`    // ID delivery date in ISO 8601 format
	RepresentativeIdExpirationDate string `json:"representative_id_expiration_date,omitempty"`  // ID expiration date in ISO 8601 format
}
