package address

import (
	"github.com/deskr/gopkg/types/country"
	"github.com/deskr/gopkg/types/geo"
)

// PostalCodeInfo holds information for post code
type PostalCodeInfo struct {
	CountryCode country.Code
	Code        string
	F1          string
	F2          string
	F3          string
	Geo         geo.Geo
}

// GetPostalCodeInfo by code
func GetPostalCodeInfo(c country.Code, code string) (p PostalCodeInfo, ok bool) {
	p, ok = postalCodeInfosByCountryCode[c][code]
	return
}
