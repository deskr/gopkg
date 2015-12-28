package address

import (
	"testing"
)

func TestNOAddress(t *testing.T) {

	adr := NOAddress{
		Recipient:   "Mike Lowrey",
		StreetName:  "Markveien",
		HouseNumber: "21",
		PostalCode:  "0554",
		Locality:    "Oslo",
	}

	pi, ok := adr.PostalCodeInfo()
	if !ok {
		t.Errorf("Failed to get post code info for: %+v", adr)
		return
	}

	if pi.F1 != "Oslo" {
		t.Errorf("Unexpected F1: %+v", pi)
		return
	}

	const adrFmt = `Mike Lowrey
Markveien 21
0554 Oslo
NORWAY`

	if adr.Format() != adrFmt {
		t.Errorf("Unexpected address format: %+v", adr.Format())
		return
	}
}
