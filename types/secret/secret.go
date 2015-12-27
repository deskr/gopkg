package secret

import (
	"errors"
	"math/rand"
	"sync"
	"time"
)

// Secret for weak to strong crypto's
type Secret string

// Grade for weak, medium, strong secrets
type Grade int

const (
	// Invalid for below weak (invalid)
	Invalid Grade = 0
	// Weak for a weak secret
	Weak Grade = 8
	// Medium for a medium secret
	Medium Grade = 16
	// Strong for a strong secret
	Strong Grade = 48
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
func NewSecret(grade Grade) Secret {
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
func (v Secret) Grade() Grade {
	l := len(v)

	// Strong secret
	if l >= int(Strong) {
		return Strong
	}

	// Medium secret
	if l >= int(Medium) {
		return Medium
	}

	// Weak secret
	if l >= int(Weak) {
		return Weak
	}

	return Invalid
}
