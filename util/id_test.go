package util

import (
	"encoding/base64"
	"strings"
	"testing"
)

func TestIDGenerate(t *testing.T) {
	g := NewIDGenerator("SomeEntity")

	id := g.Generate()

	cb, err := base64.StdEncoding.DecodeString(id)
	if err != nil {
		t.Error(err)
		return
	}

	s := string(cb)

	tid := strings.Split(s, ":")
	if len(tid) != 2 {
		t.Errorf("Failed to split ID: %s", s)
		return
	}

	if tid[0] != "SomeEntity" {
		t.Errorf("Unexpected ID decode result: %s", s)
		return
	}

	if len(tid[1]) < 6 {
		t.Errorf("Unexpected ID decode result: %s", s)
		return
	}
}
