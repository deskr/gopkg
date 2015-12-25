package types

import (
	"testing"
)

func TestValidEmails(t *testing.T) {
	t.Parallel()

	emails := [...]string{
		"mike@deskr.co",
		"json+teg@gmail.com",
		"damien.coko@vg.no",
		"pol.+tes@co.desk.com",
		"s.a.b.c@deskr.io",
		"gre+at@yahoo.tw",
		"go711a@hu31.as",
		"123@mini.br",
	}

	for _, v := range emails {
		e := Email(v)
		if !e.IsValid() {
			t.Errorf("Email is not valid: %s", v)
			return
		}
	}
}

func TestInvalidEmails(t *testing.T) {
	t.Parallel()

	emails := [...]string{
		"mike+deskr.co",
		"json+teg@gmail.",
		"damien.coko@.com",
		"pol.A+tes@co.deks.A .com.",
		"s.a.b.c@hun123",
		"123Gmail@..",
		"@gmail.com",
	}

	for _, v := range emails {
		e := Email(v)
		if e.IsValid() {
			t.Errorf("Email should not be valid: %s", v)
			return
		}
	}
}
