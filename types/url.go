package types

import (
	"net/url"
	"regexp"
	"strings"
)

// URL uniform resource locator
type URL string

var rxURL = regexp.MustCompile(`^(http(?:s)?\:\/\/[a-zA-Z0-9]+(?:(?:\.|\-)[a-zA-Z0-9]+)+(?:\:\d+)` +
	`?(?:\/[\w\-]+)*(?:\/?|\/\w+\.[a-zA-Z]{2,4}(?:\?[\w]+\=[\w\-]+)?)?(?:\&[\w]+\=[\w\-]+)*)$`)

// IsValid checks for valid username
func (v URL) IsValid() bool {
	if v == "" || len(v) >= 2083 || len(v) <= 3 || strings.HasPrefix(v.String(), ".") {
		return false
	}
	u, err := url.Parse(v.String())
	if err != nil {
		return false
	}
	if strings.HasPrefix(u.Host, ".") {
		return false
	}
	if u.Host == "" && (u.Path != "" && !strings.Contains(u.Path, ".")) {
		return false
	}
	tmp := strings.Split(u.Host, ":")
	if tmp[0] == "localhost" {
		return true
	}
	return rxURL.Match([]byte(v.String()))
}

func (v URL) String() string {
	return string(v)
}
