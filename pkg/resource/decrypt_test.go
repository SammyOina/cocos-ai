// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

package resource

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecryptFile(t *testing.T) {
	key := make([]byte, 32)
	_, err := io.ReadFull(rand.Reader, key)
	require.NoError(t, err)

	plaintext := []byte("hello world")

	// Encrypt data
	block, err := aes.NewCipher(key)
	require.NoError(t, err)

	gcm, err := cipher.NewGCM(block)
	require.NoError(t, err)

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	require.NoError(t, err)

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	tmpDir := t.TempDir()
	encryptedPath := filepath.Join(tmpDir, "encrypted.bin")
	err = os.WriteFile(encryptedPath, ciphertext, 0o644)
	require.NoError(t, err)

	t.Run("Successful decryption", func(t *testing.T) {
		decrypted, err := DecryptFile(encryptedPath, key)
		assert.NoError(t, err)
		assert.Equal(t, plaintext, decrypted)
	})

	t.Run("Invalid key size", func(t *testing.T) {
		_, err := DecryptFile(encryptedPath, key[:16])
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid key size")
	})

	t.Run("File not found", func(t *testing.T) {
		_, err := DecryptFile(filepath.Join(tmpDir, "nonexistent"), key)
		assert.Error(t, err)
	})

	t.Run("Ciphertext too short", func(t *testing.T) {
		shortPath := filepath.Join(tmpDir, "short.bin")
		err = os.WriteFile(shortPath, []byte("short"), 0o644)
		require.NoError(t, err)

		_, err = DecryptFile(shortPath, key)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "ciphertext too short")
	})

	t.Run("Decryption failed (auth error)", func(t *testing.T) {
		wrongKey := make([]byte, 32)
		_, err := io.ReadFull(rand.Reader, wrongKey)
		require.NoError(t, err)

		_, err = DecryptFile(encryptedPath, wrongKey)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "decryption failed")
	})
}
