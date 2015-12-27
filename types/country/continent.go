package country

// Continent for country
type Continent string

// IsValid checks for valid username
func (v Continent) IsValid() bool {
	if _, found := countriesByContinent[v]; !found {
		return false
	}
	return true
}

func (v Continent) String() string {
	return string(v)
}
