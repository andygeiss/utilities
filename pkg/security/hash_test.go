package security_test

import (
	"encoding/hex"
	"testing"

	"github.com/andygeiss/utilities/pkg/security"
	assert "github.com/andygeiss/utilities/pkg/testing"
)

func TestHash(t *testing.T) {
	salt := "salt"
	data := []byte("data")
	digest := security.Hash(salt, data)
	encoded := hex.EncodeToString(digest)
	assert.That("Encoded data should be equal ...", t, encoded, "856094f8ed7dc997181a936320c63be4b5bdacdbab3015ca5ab4d37bdb433760")
}
