package country

import "testing"

func TestGetCurrencySE(t *testing.T) {
	t.Parallel()

	c, ok := GetCountry(Code("SE"))
	if !ok {
		t.Error("Failed to get country SE")
		return
	}

	if c.Currency != "SEK" {
		t.Errorf("Unexpected currency: %+v", c)
		return
	}
}

func TestGetCurrencyNO(t *testing.T) {
	t.Parallel()

	c, ok := GetCountry(Code("NO"))
	if !ok {
		t.Error("Failed to get country NO")
		return
	}

	if c.Currency != "NOK" {
		t.Errorf("Unexpected currency: %+v", c)
		return
	}
}

func TestGetCurrencyUS(t *testing.T) {
	t.Parallel()

	c, ok := GetCountry(Code("US"))
	if !ok {
		t.Error("Failed to get country US")
		return
	}

	if c.Currency != "USD" {
		t.Errorf("Unexpected currency: %+v", c)
		return
	}
}

func TestGetCurrencyFR(t *testing.T) {
	t.Parallel()

	c, ok := GetCountry(Code("FR"))
	if !ok {
		t.Error("Failed to get country FR")
		return
	}

	if c.Currency != "EUR" {
		t.Errorf("Unexpected currency: %+v", c)
		return
	}
}
