package email

import (
	"regexp"
)

// Email address
type Email string

var rxEmail = regexp.MustCompile(".+@.+\\..+")

// IsValid checks for valid username
func (v Email) IsValid() bool {
	return rxEmail.Match([]byte(v.String()))
}

func (v Email) String() string {
	return string(v)
}
