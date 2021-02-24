package security

import (
	"crypto/hmac"
	"crypto/sha512"
)

// Hash HMAC-SHA512/256
func Hash(salt string, data []byte) (sum []byte) {
	hash := hmac.New(sha512.New512_256, []byte(salt))
	hash.Write(data)
	return hash.Sum(nil)
}
