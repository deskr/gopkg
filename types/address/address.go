package address

import (
	"fmt"
	"strings"

	"github.com/deskr/gopkg/types/country"
)

func trimLine(str string) string {
	return strings.TrimSpace(strings.Replace(str, "  ", " ", -1))
}

// Formatter for address
type Formatter interface {
	Format() string
}

// NOAddress for norwegian format
type NOAddress struct {
	Recipient   string `json:"recipient"`
	StreetName  string `json:"streetName"`
	HouseNumber string `json:"houseNumber"`
	PostalCode  string `json:"postalCode"`
	Locality    string `json:"locality"`
}

// Format returns a formatted address
func (a NOAddress) Format() string {
	line1 := trimLine(a.Recipient)
	line2 := trimLine(fmt.Sprintf("%s %s",
		a.StreetName,
		a.HouseNumber))
	line3 := trimLine(fmt.Sprintf("%s %s",
		a.PostalCode,
		a.Locality))
	line4 := "NORWAY"

	return fmt.Sprintf("%s\n%s\n%s\n%s",
		line1,
		line2,
		line3,
		line4)
}

// PostalCodeInfo for this address
func (a NOAddress) PostalCodeInfo() (PostalCodeInfo, bool) {
	return GetPostalCodeInfo(country.Code("NO"), a.PostalCode)
}

// USAddress for us format
type USAddress struct {
	Recipient            string `json:"recipient"`
	HouseNumber          string `json:"houseNumber"`
	StreetName           string `json:"streetName"`
	StreetType           string `json:"streetType"`
	StreetDirection      string `json:"streetDirection"`
	Building             string `json:"building"`
	Floor                string `json:"floor"`
	Apartment            string `json:"apartment"`
	ProvinceAbbreviation string `json:"provinceAbbreviation"`
	PostalCode           string `json:"postalCode"`
	Locality             string `json:"locality"`
}

// Format returns a formatted address
func (a USAddress) Format() string {
	line1 := trimLine(a.Recipient)
	line2 := trimLine(fmt.Sprintf("%s %s %s %s %s %s %s",
		a.HouseNumber,
		a.StreetName,
		a.StreetType,
		a.StreetDirection,
		a.Building,
		a.Floor,
		a.Apartment))
	line3 := trimLine(fmt.Sprintf("%s, %s %s",
		a.Locality,
		a.ProvinceAbbreviation,
		a.PostalCode))
	line4 := "UNITED STATES"

	return fmt.Sprintf("%s\n%s\n%s\n%s",
		line1,
		line2,
		line3,
		line4)
}

// PostalCodeInfo for this address
func (a USAddress) PostalCodeInfo() (PostalCodeInfo, bool) {
	return GetPostalCodeInfo(country.Code("US"), a.PostalCode)
}
