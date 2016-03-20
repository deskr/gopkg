package password

import (
	"bytes"
	"testing"
	"time"
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

func TestPasswordGenerationTime(t *testing.T) {
	clearText := "pAssWrodAbc12"

	pw := Password(clearText)

	start := time.Now()

	pw.NewHash()

	end := time.Now()

	t.Logf("Password generation time: %s", end.Sub(start).String())
}
