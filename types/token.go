package types

import (
	"crypto/sha256"
	"fmt"
)

// Token for use with urls
type Token string

// NewToken generates a new token
func NewToken() Token {
	return Token(generateToken())
}

func generateToken() string {
	s := fmt.Sprintf("%s%s", NewUUID().String(),
		NewSecret(SecretMedium).String())
	h := sha256.New224()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}

func (v Token) String() string {
	return string(v)
}
