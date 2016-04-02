package language

import "strings"

// Language ISO 639-1
type Language struct {
	Code       Code   `json:"code"`
	Name       string `json:"name"`
	NativeName string `json:"native_name"`
}

// Code ISO 639-1
type Code string

// IsValid checks for valid username
func (v Code) IsValid() bool {
	return true
}

func (v Code) String() string {
	return string(v)
}

// Parse tries to find a language
func Parse(s string) (Language, bool) {
	val := strings.TrimSpace(strings.ToLower(s))

	if v, ok := languagesByCode[Code(val)]; ok {
		return v, true
	}

	if strings.Contains(val, "-") {
		for _, v := range strings.Split(val, "-") {
			if l, ok := languagesByCode[Code(v)]; ok {
				return l, true
			}
		}
	}

	return Language{}, false
}
