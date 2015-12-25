package types

import (
	"regexp"
)

// IPAddress IP address
type IPAddress string

var rxIP = regexp.MustCompile(`^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`)

// IsValid checks for valid username
func (v IPAddress) IsValid() bool {
	return rxIP.Match([]byte(v.String()))
}

func (v IPAddress) String() string {
	return string(v)
}
