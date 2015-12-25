package types

// Username for user
type Username string

// IsValid checks for valid username
func (v Username) IsValid() bool {
	return Email(v.String()).IsValid()
}

func (v Username) String() string {
	return string(v)
}
