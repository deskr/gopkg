package currency

import "testing"

func TestGetCurrencySE(t *testing.T) {
	t.Parallel()

	c, ok := GetCurrency(Code("SEK"))
	if !ok {
		t.Error("Failed to get currency SEK")
		return
	}

	if c.Name != "Swedish Krona" {
		t.Errorf("Unexpected currency: %+v", c)
		return
	}
}

func TestGetCurrencyNO(t *testing.T) {
	t.Parallel()

	c, ok := GetCurrency(Code("NOK"))
	if !ok {
		t.Error("Failed to get currency NOK")
		return
	}

	if c.Name != "Norwegian Krone" {
		t.Errorf("Unexpected currency: %+v", c)
		return
	}
}

func TestGetCurrencyUS(t *testing.T) {
	t.Parallel()

	c, ok := GetCurrency(Code("USD"))
	if !ok {
		t.Error("Failed to get currency USD")
		return
	}

	if c.Name != "US Dollar" {
		t.Errorf("Unexpected currency: %+v", c)
		return
	}
}

func TestGetCurrencyFR(t *testing.T) {
	t.Parallel()

	c, ok := GetCurrency(Code("EUR"))
	if !ok {
		t.Error("Failed to get currency EUR")
		return
	}

	if c.Name != "Euro" {
		t.Errorf("Unexpected currency: %+v", c)
		return
	}
}
