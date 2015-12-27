package ip

import (
	"regexp"
)

// Address IP address
type Address string

var rxIP = regexp.MustCompile(`^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`)

// IsValid checks for valid username
func (v Address) IsValid() bool {
	return rxIP.Match([]byte(v.String()))
}

func (v Address) String() string {
	return string(v)
}
