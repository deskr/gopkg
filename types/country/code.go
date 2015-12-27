package country

// Code ISO 3166-1 alpha 2
type Code string

// IsValid checks for valid username
func (v Code) IsValid() bool {
	if _, found := countriesByCode[v]; !found {
		return false
	}
	return true
}

func (v Code) String() string {
	return string(v)
}
