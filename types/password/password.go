package password

import (
	"bytes"
	"crypto/sha1"
	"io"

	"github.com/deskr/gopkg/types/secret"

	"github.com/magical/argon2"
)

// Password for user
type Password string

// Hash for the hashed password
type Hash struct {
	Hash []byte
	Salt []byte
}

// MinimumLength is the minimum password length
const MinimumLength = 6

// generateSalt generates a random strong salt to use in hashing secret stuff
func generateSalt() []byte {
	h := sha1.New()
	io.WriteString(h, secret.NewSecret(secret.Medium).String())
	return h.Sum(nil)
}

// generateHash returns hash
func generateHash(str string, salt []byte) []byte {
	key, err := argon2.Key([]byte(str), salt, 3, 1, 8192, 64)
	if err != nil {
		panic(err)
	}
	return key
}

// IsValid checks for valid password
func (v Password) IsValid() bool {
	if len(v.String()) < MinimumLength {
		return false
	}
	return true
}

func (v Password) String() string {
	return string(v)
}

// NewHash creates a new hash variant (including salt) of the clear text password
// this method will return a new value each time called
func (v Password) NewHash() Hash {
	salt := generateSalt()
	hash := generateHash(string(v), salt)

	return Hash{
		Hash: hash,
		Salt: salt,
	}
}

// Validate checks the clear-text password against the stored hash password
func (v Hash) Validate(password Password) bool {
	return bytes.Equal(generateHash(password.String(), v.Salt), v.Hash)
}
