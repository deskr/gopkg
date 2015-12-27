package currency

// Code for currency
type Code string

// IsValid checks for valid username
func (v Code) IsValid() bool {
	if _, found := currenciesByCode[v]; !found {
		return false
	}
	return true
}

func (v Code) String() string {
	return string(v)
}
