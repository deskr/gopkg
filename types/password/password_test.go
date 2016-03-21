package password

import (
	"bytes"
	"testing"
)

func TestPasswordValidate(t *testing.T) {
	t.Parallel()

	clearText := "pAssWrodAbc12"

	pw := Password(clearText)

	ph := pw.NewHash()

	ok := ph.Validate(pw)

	if !ok {
		t.Error("Password validation failed")
	}
}

func TestPasswordNewHash(t *testing.T) {
	t.Parallel()

	clearText := "pAssWrodAbc12"

	pw := Password(clearText)

	ph1 := pw.NewHash()
	ph2 := pw.NewHash()

	if bytes.Equal(ph1.Hash, ph2.Hash) {
		t.Error("Password NewHash generated the same hash output twice")
	}

	if bytes.Equal(ph1.Salt, ph2.Salt) {
		t.Error("Password NewHash generated the same salt output twice")
	}
}

func createPasswords(b *testing.B, n int) {
	var ps []Hash

	for i := 0; i < n; i++ {
		clearText := "pAssWrodAbc12"
		pw := Password(clearText)
		h := pw.NewHash()

		ps = append(ps, h)
	}
}

func BenchmarkPasswordGeneration1(b *testing.B) {
	createPasswords(b, 1)
}

func BenchmarkPasswordGeneration100(b *testing.B) {
	createPasswords(b, 100)
}

func BenchmarkPasswordGeneration1000(b *testing.B) {
	createPasswords(b, 1000)
}
