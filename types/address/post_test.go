package address

import (
	"testing"

	"github.com/deskr/gopkg/types/country"
)

func TestPostalCodeInfo0554OsloNO(t *testing.T) {
	t.Parallel()

	p, ok := GetPostalCodeInfo(country.Code("NO"), "0554")
	if !ok {
		t.Errorf("Failed to get post code info for NO")
		return
	}

	if p.F1 != "Oslo" {
		t.Errorf("Unexpected F1: %+v", p)
		return
	}
}

func TestPostalCodeInfo2685GarmoNO(t *testing.T) {
	t.Parallel()

	p, ok := GetPostalCodeInfo(country.Code("NO"), "2685")
	if !ok {
		t.Errorf("Failed to get post code info for NO")
		return
	}

	if p.F1 != "Garmo" {
		t.Errorf("Unexpected F1: %+v", p)
		return
	}

	if p.F2 != "Oppland" {
		t.Errorf("Unexpected F2: %+v", p)
	}

	if p.F3 != "Lom" {
		t.Errorf("Unexpected F3: %+v", p)
	}
}

func TestPostalCodeInfo29403CharlestonUS(t *testing.T) {
	t.Parallel()

	p, ok := GetPostalCodeInfo(country.Code("US"), "29403")
	if !ok {
		t.Errorf("Failed to get post code info for US")
		return
	}

	if p.F1 != "Charleston" {
		t.Errorf("Unexpected F1: %+v", p)
		return
	}
}

func TestPostalCodeInfo13520000SaoPedroBR(t *testing.T) {
	t.Parallel()

	p, ok := GetPostalCodeInfo(country.Code("BR"), "13520-000")
	if !ok {
		t.Errorf("Failed to get post code info for BR")
		return
	}

	if p.F1 != "São Pedro" {
		t.Errorf("Unexpected F1: %+v", p)
		return
	}
}

func TestPostalCodeInfo301007МалаховоRU(t *testing.T) {
	t.Parallel()

	p, ok := GetPostalCodeInfo(country.Code("RU"), "301007")
	if !ok {
		t.Errorf("Failed to get post code info for RU")
		return
	}

	if p.F1 != "Малахово" {
		t.Errorf("Unexpected F1: %+v", p)
		return
	}
}
