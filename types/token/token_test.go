package token

import (
	"testing"
)

func TestToken(t *testing.T) {
	t.Parallel()

	tokens := make(map[string]struct{})
	var v struct{}

	for i := 0; i < 10000; i++ {
		token := NewToken()

		if v, exists := tokens[token.String()]; exists {
			t.Errorf("Token not unique, previously generated: %s", v)
			return
		}
		tokens[token.String()] = v
	}
}
