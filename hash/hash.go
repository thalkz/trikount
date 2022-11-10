package hash

import (
	"encoding/base64"
	"math/rand"
	"time"
)

func NewHash(size int) string {
	rand.Seed(time.Now().Unix())
	bytes := make([]byte, size)
	for i := range bytes {
		bytes[i] = byte(rand.Intn(128))
	}
	return base64.RawURLEncoding.EncodeToString(bytes)[:size]
}
