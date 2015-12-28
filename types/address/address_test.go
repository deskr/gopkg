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
		t.Errorf("Unexpected address format.\nGot:\n\"%+v\"\n\nExpected:\n\"%+v\"", adr.Format(), adrFmt)
		return
	}
}

func TestUSAddress(t *testing.T) {

	adr := USAddress{
		Recipient:            "Jeremy Martinson",
		HouseNumber:          "455",
		StreetName:           "Larkspur Dr.",
		Locality:             "San Jose",
		ProvinceAbbreviation: "CA",
		PostalCode:           "95111",
	}

	pi, ok := adr.PostalCodeInfo()
	if !ok {
		t.Errorf("Failed to get post code info for: %+v", adr)
		return
	}

	if pi.F1 != "San Jose" {
		t.Errorf("Unexpected F1: %+v", pi)
		return
	}

	if pi.F2 != "California" {
		t.Errorf("Unexpected F2: %+v", pi)
		return
	}

	if pi.F3 != "Santa Clara" {
		t.Errorf("Unexpected F3: %+v", pi)
		return
	}

	const adrFmt = `Jeremy Martinson
455 Larkspur Dr.
San Jose, CA 95111
UNITED STATES`

	if adr.Format() != adrFmt {
		t.Errorf("Unexpected address format.\nGot:\n\"%+v\"\n\nExpected:\n\"%+v\"", adr.Format(), adrFmt)
		return
	}
}
