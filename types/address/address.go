package address

import (
	"fmt"
	"strings"

	"github.com/deskr/gopkg/types/country"
)

// NOAddress for norwegian format
// RECIPIENT
// [STREET_NAME HOUSE_NUMBER]
// POSTAL_CODE LOCALITY
// NORWAY
type NOAddress struct {
	Recipient   string
	StreetName  string
	HouseNumber string
	PostalCode  string
	Locality    string
}

// Format returns a formatted address
func (a NOAddress) Format() string {
	return strings.TrimSpace(strings.Replace(fmt.Sprintf(`%s
%s %s
%s %s
%s`,
		a.Recipient,
		a.StreetName,
		a.HouseNumber,
		a.PostalCode,
		a.Locality,
		"NORWAY"), "  ", " ", -1))
}

// PostalCodeInfo for this address
func (a NOAddress) PostalCodeInfo() (PostalCodeInfo, bool) {
	return GetPostalCodeInfo(country.Code("NO"), a.PostalCode)
}

// USAddress for us format
// RECIPIENT
// HOUSE_NUMBER STREET_NAME [STREET_TYPE] [STREET_DIRECTION] [BUILDING] [FLOOR] [APARTMENT]
// LOCALITY PROVINCE_ABBREVIATION POSTAL_CODE
// UNITED STATES
type USAddress struct {
	Recipient            string
	HouseNumber          string
	StreetName           string
	StreetType           string
	StreetDirection      string
	Building             string
	Floor                string
	Apartment            string
	ProvinceAbbreviation string
	PostalCode           string
	Locality             string
}

// Format returns a formatted address
func (a USAddress) Format() string {
	return strings.TrimSpace(strings.Replace(fmt.Sprintf(`%s
%s %s %s %s %s %s %s
%s %s %s
%s`,
		a.Recipient,
		a.HouseNumber,
		a.StreetName,
		a.StreetType,
		a.StreetDirection,
		a.Building,
		a.Floor,
		a.Apartment,
		a.Locality,
		a.ProvinceAbbreviation,
		a.PostalCode,
		"UNITED STATES"), "  ", " ", -1))
}

// PostalCodeInfo for this address
func (a USAddress) PostalCodeInfo() (PostalCodeInfo, bool) {
	return GetPostalCodeInfo(country.Code("NO"), a.PostalCode)
}
