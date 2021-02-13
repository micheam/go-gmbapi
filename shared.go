package gmbapi

// OrganizationInfo is an Additional Info stored for an organization.
type OrganizationInfo struct {
	RegisteredDomain *string        `json:"registeredDomain"`
	PostalAddress    *PostalAddress `json:"postalAddress"`
	PhoneNumber      *string        `json:"phoneNumber"`
}

// PostalAddress Represents a postal address,
// e.g. for postal delivery or payments addresses.
type PostalAddress struct {
	Revision           int64    `json:"revision"`
	RegionCode         string   `json:"regionCode"`
	LanguageCode       string   `json:"languageCode"`
	PostalCode         string   `json:"postalCode"`
	SortingCode        string   `json:"sortingCode"`
	AdministrativeArea string   `json:"administrativeArea"`
	Locality           string   `json:"locality"`
	Sublocality        string   `json:"sublocality"`
	AddressLines       []string `json:"addressLines"`
	Recipients         []string `json:"recipients"`
	Organization       string   `json:"organization"`
}
