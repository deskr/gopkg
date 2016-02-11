package country

import "github.com/deskr/gopkg/types/language"
import "github.com/deskr/gopkg/types/currency"

// Country contains info about a country
type Country struct {
	Code          Code            `json:"code"`
	Name          string          `json:"name"`
	Capital       string          `json:"capital"`
	Continent     string          `json:"continent"`
	ContinentName string          `json:"continentName"`
	CurrencyCode  currency.Code   `json:"currencyCode"`
	LanguageCodes []language.Code `json:"languageCodes"`
	VatRates      VatRates        `json:"vatRates"`
}

// VatRates for country
type VatRates struct {
	Standard float64 `json:"standard"`
	Reduced1 float64 `json:"reduced1"`
	Reduced2 float64 `json:"reduced2"`
}

// Languages gets the country languages
func (c Country) Languages() []language.Language {
	var list []language.Language
	for _, v := range c.LanguageCodes {
		l, _ := language.Parse(v.String())
		list = append(list, l)
	}
	return list
}

// Currency gets the country currency
func (c Country) Currency() currency.Currency {
	v, _ := currency.GetCurrency(c.CurrencyCode)
	return v
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
