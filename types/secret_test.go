package types

import (
	"testing"
)

func TestSecret(t *testing.T) {
	t.Parallel()

	for i := 0; i < 100; i++ {
		s := NewSecret(SecretStrong)
		if s.Grade() != SecretStrong {
			t.Errorf("Secret was not strong: %s", s)
		}
	}

	for i := 0; i < 100; i++ {
		s := NewSecret(SecretMedium)
		if s.Grade() != SecretMedium {
			t.Errorf("Secret was not medium: %s", s)
		}
	}

	for i := 0; i < 100; i++ {
		s := NewSecret(SecretWeak)
		if s.Grade() != SecretWeak {
			t.Errorf("Secret was not weak: %s", s)
		}
	}

	for i := 0; i < 100; i++ {
		s := NewSecret(SecretInvalid)
		if s.Grade() != SecretInvalid {
			t.Errorf("Secret was not invalid: %s", s)
		}
	}
}
