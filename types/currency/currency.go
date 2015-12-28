package currency

// Currency for country
type Currency struct {
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	NativeSymbol  string `json:"nativeSymbol"`
	PluralName    string `json:"pluralName"`
	DecimalDigits int    `json:"decimalDigits"`
	Code          Code   `json:"code"`
}

// GetCurrency gets a currency by a currency code
func GetCurrency(code Code) (c Currency, ok bool) {
	ok = false
	if !code.IsValid() {
		return
	}

	c, ok = currenciesByCode[code]
	return
}
