package address

import (
	"github.com/deskr/gopkg/types/country"
	"github.com/deskr/gopkg/types/geo"
)

// PostalCodeInfo holds information for post code
type PostalCodeInfo struct {
	CountryCode country.Code `json:"countryCode"`
	Code        string       `json:"code"`
	F1          string       `json:"f1"`
	F2          string       `json:"f2"`
	F3          string       `json:"f3"`
	Geo         geo.Geo      `json:"geo"`
}

// GetPostalCodeInfo by code
func GetPostalCodeInfo(c country.Code, code string) (p PostalCodeInfo, ok bool) {
	p, ok = postalCodeInfosByCountryCode[c][code]
	return
}
