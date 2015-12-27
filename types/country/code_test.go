package country

import "testing"

func TestGetCodeBR(t *testing.T) {
	t.Parallel()

	_, ok := GetCountry(Code("BR"))
	if !ok {
		t.Error("Failed to get country BR")
		return
	}
}

func TestGetCodeIT(t *testing.T) {
	t.Parallel()

	_, ok := GetCountry(Code("IT"))
	if !ok {
		t.Error("Failed to get country IT")
		return
	}
}
