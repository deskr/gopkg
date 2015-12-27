package currency

import "testing"

func TestGetCodeBR(t *testing.T) {
	t.Parallel()

	_, ok := GetCurrency(Code("NOK"))
	if !ok {
		t.Error("Failed to get currency NOK")
		return
	}
}

func TestGetCodeIT(t *testing.T) {
	t.Parallel()

	_, ok := GetCurrency(Code("USD"))
	if !ok {
		t.Error("Failed to get currency USD")
		return
	}
}
