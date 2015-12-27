package password

import (
	"bytes"
	"crypto/sha1"
	"io"

	"github.com/deskr/gopkg/types/secret"

	"golang.org/x/crypto/pbkdf2"
)

// Password for user
type Password string

// Hash is the hashed version of the clear-text password
// The salt used in creating the hash is kept along the password.
type Hash struct {
	Hash []byte
	Salt []byte
}

// MinimumLength is the minimum password length
const MinimumLength = 6

const (
	hashIterations = 4096
	keyLength      = 32
)

// generateSalt generates a random strong salt to use in hashing secret stuff
func generateSalt() []byte {
	h := sha1.New()
	io.WriteString(h, secret.NewSecret(secret.Medium).String())
	return h.Sum(nil)
}

// generateHash returns hash
func generateHash(str string, salt []byte) []byte {
	return pbkdf2.Key([]byte(string(salt)+str), salt, hashIterations, keyLength, sha1.New)
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
	hash := generateHash(v.String(), salt)
	return Hash{
		Hash: hash,
		Salt: salt,
	}
}

// Validate checks the clear-text password against the stored hash password
func (v Hash) Validate(password Password) bool {
	return bytes.Equal(generateHash(password.String(), v.Salt), v.Hash)
}
