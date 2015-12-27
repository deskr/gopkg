package currency

// Currency for country
type Currency struct {
	Name          string
	Symbol        string
	NativeSymbol  string
	PluralName    string
	DecimalDigits int
	Code          Code
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
