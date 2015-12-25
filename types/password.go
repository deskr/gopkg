package types

import (
	"bytes"
	"crypto/sha1"
	"io"

	"golang.org/x/crypto/pbkdf2"
)

// Password for user
type Password string

// PasswordHash is the hashed version of the clear-text password
// The salt used in creating the hash is kept along the password.
type PasswordHash struct {
	Hash []byte
	Salt []byte
}

// MinimumPasswordLength is the minimum password length
const MinimumPasswordLength = 6

const (
	hashIterations = 4096 // DO NOT CHANGE - will invalidate all current passwords
	keyLength      = 32   // DO NOT CHANGE
)

// generateSalt generates a random strong salt to use in hashing secret stuff
func generateSalt() []byte {
	h := sha1.New()
	io.WriteString(h, NewSecret(SecretMedium).String())
	return h.Sum(nil)
}

// generateHash returns hash
func generateHash(str string, salt []byte) []byte {
	return pbkdf2.Key([]byte(string(salt)+str), salt, hashIterations, keyLength, sha1.New)
}

// IsValid checks for valid password
func (v Password) IsValid() bool {
	if len(v.String()) < MinimumPasswordLength {
		return false
	}
	return true
}

func (v Password) String() string {
	return string(v)
}

// NewHash creates a new hash variant (including salt) of the clear text password
// this method will return a new value each time called
func (v Password) NewHash() PasswordHash {
	salt := generateSalt()
	hash := generateHash(v.String(), salt)
	return PasswordHash{
		Hash: hash,
		Salt: salt,
	}
}

// Validate checks the clear-text password against the stored hash password
func (v PasswordHash) Validate(password Password) bool {
	return bytes.Equal(generateHash(password.String(), v.Salt), v.Hash)
}
