package country

// Currency for country
type Currency string

// IsValid checks for valid username
func (v Currency) IsValid() bool {
	if _, found := countriesByCurrency[v]; !found {
		return false
	}
	return true
}

func (v Currency) String() string {
	return string(v)
}
