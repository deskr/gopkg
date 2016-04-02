package address

import (
	"testing"
)

func TestNOAddress(t *testing.T) {

	adr := NOAddress{
		Recipient:  "Mike Lowrey",
		Street:     "Markveien 21",
		PostalCode: "0554",
		Locality:   "Oslo",
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
		Recipient:  "Jeremy Martinson",
		Street:     "455 Larkspur Dr.",
		Locality:   "San Jose",
		StateCode:  "CA",
		PostalCode: "95111",
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
