package country

import "testing"

func TestGetContinentNL(t *testing.T) {
	t.Parallel()

	c, ok := GetCountry(Code("NL"))
	if !ok {
		t.Error("Failed to get country NL")
		return
	}

	if c.ContientName != "Europe" {
		t.Errorf("Unexpected continent name: %+v", c)
		return
	}

	if c.Continent != "EU" {
		t.Errorf("Unexpected continent: %+v", c)
		return
	}
}

func TestGetContinentBR(t *testing.T) {
	t.Parallel()

	c, ok := GetCountry(Code("BR"))
	if !ok {
		t.Error("Failed to get country BR")
		return
	}

	if c.ContientName != "South America" {
		t.Errorf("Unexpected continent name: %+v", c)
		return
	}

	if c.Continent != "SA" {
		t.Errorf("Unexpected continent: %+v", c)
		return
	}
}
