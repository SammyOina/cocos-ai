// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

package resource

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"testing"
)

func TestDecryptData(t *testing.T) {
	key := make([]byte, 32) // AES-256
	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		t.Fatalf("failed to generate key: %v", err)
	}

	plaintext := []byte("secret information")

	// Encrypt to formulate test vector
	block, err := aes.NewCipher(key)
	if err != nil {
		t.Fatalf("failed to create cipher: %v", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		t.Fatalf("failed to create gcm: %v", err)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		t.Fatalf("failed to generate nonce: %v", err)
	}

	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)

	// In CoCo format, they want nonce prepended
	encryptedData := append(nonce, ciphertext...)

	// Test DecryptData successful
	decrypted, err := DecryptData(encryptedData, key)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if string(decrypted) != string(plaintext) {
		t.Fatalf("expected %q, got %q", string(plaintext), string(decrypted))
	}

	// Test invalid key size
	_, err = DecryptData(encryptedData, make([]byte, 16))
	if err == nil {
		t.Fatalf("expected error for invalid key size")
	}

	// Test too short data
	_, err = DecryptData(make([]byte, 10), key)
	if err == nil {
		t.Fatalf("expected error for too short data")
	}

	// Test corrupted data
	encryptedData[15] ^= 1 // Flip a bit
	_, err = DecryptData(encryptedData, key)
	if err == nil {
		t.Fatalf("expected error for corrupted ciphertext")
	}
}
