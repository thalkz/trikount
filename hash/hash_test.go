package hash_test

import (
	"testing"

	"github.com/thalkz/trikount/hash"
)

func TestNewProjectId(T *testing.T) {
	for i := 0; i < 100; i++ {
		id := hash.NewHash(i)
		if len(id) != i {
			T.Fatalf("invalid hash: %v", id)
		}
	}
}
