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
	Recipient  string `json:"recipient"`
	Street     string `json:"street"`
	PostalCode string `json:"postalCode"`
	Locality   string `json:"locality"`
}

// Format returns a formatted address
func (a NOAddress) Format() string {
	return fmt.Sprintf("%s\n%s\n%s\n%s",
		trimLine(a.Recipient),
		trimLine(a.Street),
		trimLine(fmt.Sprintf("%s %s",
			a.PostalCode,
			a.Locality)),
		"NORWAY")
}

// PostalCodeInfo for this address
func (a NOAddress) PostalCodeInfo() (PostalCodeInfo, bool) {
	return GetPostalCodeInfo(country.Code("NO"), a.PostalCode)
}

// USAddress for us format
type USAddress struct {
	Recipient  string `json:"recipient"`
	Street     string `json:"street"`
	Locality   string `json:"locality"`
	StateCode  string `json:"stateCode"`
	PostalCode string `json:"postalCode"`
}

// Format returns a formatted address
func (a USAddress) Format() string {
	return fmt.Sprintf("%s\n%s\n%s\n%s",
		trimLine(a.Recipient),
		trimLine(a.Street),
		trimLine(fmt.Sprintf("%s, %s %s",
			a.Locality,
			a.StateCode,
			a.PostalCode)),
		"UNITED STATES")
}

// PostalCodeInfo for this address
func (a USAddress) PostalCodeInfo() (PostalCodeInfo, bool) {
	return GetPostalCodeInfo(country.Code("US"), a.PostalCode)
}
