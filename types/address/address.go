package address

import (
	"fmt"
	"strings"
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
	PostalCode string `json:"postal_code"`
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

// USAddress for us format
type USAddress struct {
	Recipient  string `json:"recipient"`
	Street     string `json:"street"`
	Locality   string `json:"locality"`
	StateCode  string `json:"state_code"`
	PostalCode string `json:"postal_code"`
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
