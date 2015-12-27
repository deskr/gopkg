package country

import "github.com/deskr/gopkg/types/language"

// Country contains info about a country
type Country struct {
	Code         Code
	Name         string
	Capital      string
	Currency     Currency
	Continent    Continent
	ContientName string
	Languages    []language.Language
}

// GetCountry gets a country by a country code
func GetCountry(code Code) (c Country, ok bool) {
	ok = false
	if !code.IsValid() {
		return
	}

	c, ok = countriesByCode[code]
	return
}
