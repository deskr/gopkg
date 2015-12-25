package types

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

// Secret for weak to strong crypto's
type Secret string

// SecretGrade for weak, medium, strong secrets
type SecretGrade int

const (
	// SecretInvalid for below weak (invalid)
	SecretInvalid SecretGrade = 0
	// SecretWeak for a weak secret
	SecretWeak SecretGrade = 8
	// SecretMedium for a medium secret
	SecretMedium SecretGrade = 16
	// SecretStrong for a strong secret
	SecretStrong SecretGrade = 48
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// ErrInvalidSecret for when a secret is not valid (under weak or above insane strong)
var ErrInvalidSecret = errors.New("Invalid secret")

var randSrc = rand.NewSource(time.Now().UnixNano())
var mutex = sync.Mutex{}

// NewSecret generates a new secret
func NewSecret(grade SecretGrade) Secret {
	return Secret(randString(int(grade)))
}

// randString generates a random string of n length
func randString(n int) string {
	b := make([]byte, n)
	// A randSrc.Int63() generates 63 random bits, enough for letterIdxMax characters!
	mutex.Lock()
	for i, cache, remain := n-1, randSrc.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	mutex.Unlock()
	return string(b)
}

func (v Secret) String() string {
	return string(v)
}

// Grade gets the secret grade weak, medium or strong
func (v Secret) Grade() SecretGrade {
	l := len(v)

	// Strong secret
	if l >= int(SecretStrong) {
		return SecretStrong
	}

	// Medium secret
	if l >= int(SecretMedium) {
		return SecretMedium
	}

	// Weak secret
	if l >= int(SecretWeak) {
		return SecretWeak
	}

	return SecretInvalid
}
