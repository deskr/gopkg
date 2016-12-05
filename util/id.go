package util

import (
	"encoding/base64"
	"fmt"
	"hash/fnv"

	"github.com/ventu-io/go-shortid"
)

// IDGenerator for generating id's similiar to
// the id's used in GraphQL type:id (base64 encoded)
type IDGenerator struct {
	sid *shortid.Shortid
	typ string
}

// NewIDGenerator creates a new generator
func NewIDGenerator(typeID string) *IDGenerator {
	h := fnv.New64a()
	h.Write([]byte(typeID))
	s := h.Sum64()

	return &IDGenerator{
		sid: shortid.MustNew(1, shortid.DEFAULT_ABC, s),
		typ: typeID,
	}
}

// Generate an id
func (g IDGenerator) Generate() string {
	id := g.sid.MustGenerate()
	data := []byte(fmt.Sprintf("%s:%s", g.typ, id))
	return base64.RawURLEncoding.EncodeToString(data)
}
