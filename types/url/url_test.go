package url

import (
	"testing"
)

func TestValidURLS(t *testing.T) {
	t.Parallel()

	urls := [...]string{
		"http://google.com",
		"http://vg.no",
		"http://deskr.co",
		"https://deskr.co",
		"https://www.deskr.co:432",
		"https://sub.deskr.co",
		"https://sub.deskr.co.uk",
		"http://127.0.0.1:3001",
		"http://localhost:3001",
		"http://localhost",
	}

	for _, v := range urls {
		e := URL(v)
		if !e.IsValid() {
			t.Errorf("Email is not valid: %s", v)
			return
		}
	}
}

func TestInvalidURLS(t *testing.T) {
	t.Parallel()

	urls := [...]string{
		"http:/invalid.com",
		"htps::/nsa.us.gov",
		"http::/aften.com",
		"ttps://google.com",
		"http:/|google",
		"http://.com",
		"http:/\\test.com",
		"http//:valigs.com",
		"http:\\/\\/google.com",
	}

	for _, v := range urls {
		e := URL(v)
		if e.IsValid() {
			t.Errorf("URL should not be valid: %s", v)
			return
		}
	}
}
