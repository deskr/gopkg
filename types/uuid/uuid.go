package uuid

import (
	"regexp"

	goUUID "github.com/satori/go.uuid"
)

// UUID for id
type UUID string

// NewUUID generates a new UUID
func NewUUID() UUID {
	return UUID(goUUID.NewV4().String())
}

// IsValid checks for valid username
func (v UUID) IsValid() bool {
	if v.String() == "" {
		return false
	}
	if matched := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[a-fA-F0-9]` +
		`{4}-[a-fA-F0-9]{4}-[a-fA-F0-9]{12}$`).
		Match([]byte(v)); !matched {
		return false
	}
	return true
}

func (v UUID) String() string {
	return string(v)
}
