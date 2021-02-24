package security_test

import (
	"testing"

	"github.com/andygeiss/utilities/security"
	assert "github.com/andygeiss/utilities/testing"
)

func TestDecrypt(t *testing.T) {
	plaintext := []byte("asdf 1234")
	key := security.NewKey256()
	ciphertext, _ := security.Encrypt(plaintext, key)
	decrypted, err := security.Decrypt(ciphertext, key)
	assert.That("Decrypt should return without an error", t, err, nil)
	assert.That("Decrypted ciphertext should be equal with plaintext", t, decrypted, plaintext)
}

func TestEncrypt(t *testing.T) {
	plaintext := []byte("asdf 1234")
	key := security.NewKey256()
	ciphertext, err := security.Encrypt(plaintext, key)
	assert.That("Encrypt should return without an error", t, err, nil)
	assert.That("Ciphertext should be not empty", t, len(ciphertext) > 0, true)
}
