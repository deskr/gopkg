package country

import "testing"

func TestGetCountryUS(t *testing.T) {
	t.Parallel()

	c, ok := GetCountry(Code("US"))
	if !ok {
		t.Error("Failed to get country US")
		return
	}

	if c.Name != "United States" {
		t.Errorf("Unexpected name: %+v", c)
		return
	}

	found := false
	for _, v := range c.Languages() {
		if v.Name == "English" {
			found = true
		}
	}

	if !found {
		t.Errorf("Unexpected languages: %+v", c)
		return
	}
}

func TestGetCountrySweden(t *testing.T) {
	t.Parallel()

	c, ok := GetCountry(Code("SE"))
	if !ok {
		t.Error("Failed to get country SE")
		return
	}

	if c.Name != "Sweden" {
		t.Errorf("Unexpected name: %+v", c)
		return
	}

	found := false
	for _, v := range c.Languages() {
		if v.NativeName == "Svenska" {
			found = true
		}
	}

	if !found {
		t.Errorf("Unexpected languages: %+v", c)
		return
	}
}
