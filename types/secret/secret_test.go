package secret

import (
	"testing"
)

func TestSecret(t *testing.T) {
	t.Parallel()

	for i := 0; i < 100; i++ {
		s := NewSecret(Strong)
		if s.Grade() != Strong {
			t.Errorf("Secret was not strong: %s", s)
		}
	}

	for i := 0; i < 100; i++ {
		s := NewSecret(Medium)
		if s.Grade() != Medium {
			t.Errorf("Secret was not medium: %s", s)
		}
	}

	for i := 0; i < 100; i++ {
		s := NewSecret(Weak)
		if s.Grade() != Weak {
			t.Errorf("Secret was not weak: %s", s)
		}
	}

	for i := 0; i < 100; i++ {
		s := NewSecret(Invalid)
		if s.Grade() != Invalid {
			t.Errorf("Secret was not invalid: %s", s)
		}
	}
}
