// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

package resource

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"os"
)

const (
	// AES-GCM nonce size in bytes.
	aesGCMNonceSize = 12
	// AES-256 key size in bytes.
	aes256KeySize = 32
)

// DecryptFile decrypts an AES-256-GCM encrypted file in place.
// The encrypted file format is: nonce (12 bytes) || ciphertext+tag.
// The key must be exactly 32 bytes (AES-256).
func DecryptFile(encryptedPath string, key []byte) ([]byte, error) {
	if len(key) != aes256KeySize {
		return nil, fmt.Errorf("invalid key size: expected %d bytes, got %d", aes256KeySize, len(key))
	}

	ciphertext, err := os.ReadFile(encryptedPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read encrypted file: %w", err)
	}

	return DecryptData(ciphertext, key)
}

// DecryptData decrypts AES-256-GCM encrypted data.
// The encrypted data format is: nonce (12 bytes) || ciphertext+tag.
func DecryptData(ciphertext, key []byte) ([]byte, error) {
	if len(key) != aes256KeySize {
		return nil, fmt.Errorf("invalid key size: expected %d bytes, got %d", aes256KeySize, len(key))
	}

	if len(ciphertext) < aesGCMNonceSize {
		return nil, fmt.Errorf("ciphertext too short: expected at least %d bytes for nonce, got %d", aesGCMNonceSize, len(ciphertext))
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("failed to create AES cipher: %w", err)
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("failed to create GCM cipher: %w", err)
	}

	nonce := ciphertext[:aesGCMNonceSize]
	encData := ciphertext[aesGCMNonceSize:]

	plaintext, err := gcm.Open(nil, nonce, encData, nil)
	if err != nil {
		return nil, fmt.Errorf("decryption failed (authentication error): %w", err)
	}

	return plaintext, nil
}
