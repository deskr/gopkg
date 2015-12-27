package language

import "testing"

func TestParseLanguage(t *testing.T) {
	t.Parallel()

	c, ok := Parse("ar-AE")
	if !ok {
		t.Error("Failed to parse language ar-AE")
		return
	}

	if c.Name != "Arabic" {
		t.Errorf("Unexpected language name: %+v", c)
		return
	}
}
