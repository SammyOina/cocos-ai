// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

package resource

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseS3URL(t *testing.T) {
	tests := []struct {
		url    string
		bucket string
		key    string
		err    bool
	}{
		{"s3://my-bucket/my-key", "my-bucket", "my-key", false},
		{"s3://my-bucket/path/to/my-key", "my-bucket", "path/to/my-key", false},
		{"s3://my-bucket/", "", "", true},
		{"s3://", "", "", true},
		{"http://my-bucket/my-key", "", "", true},
		{"s3://my-bucket", "", "", true},
	}

	for _, tt := range tests {
		bucket, key, err := parseS3URL(tt.url)
		if tt.err {
			if err == nil {
				t.Errorf("expected error for %s, got nil", tt.url)
			}
		} else {
			if err != nil {
				t.Errorf("expected no error for %s, got %v", tt.url, err)
			}
			if bucket != tt.bucket {
				t.Errorf("expected bucket %s, got %s", tt.bucket, bucket)
			}
			if key != tt.key {
				t.Errorf("expected key %s, got %s", tt.key, key)
			}
		}
	}
}

func TestParseGCSURL(t *testing.T) {
	tests := []struct {
		url    string
		bucket string
		key    string
		err    bool
	}{
		{"gs://my-bucket/my-key", "my-bucket", "my-key", false},
		{"gs://my-bucket/path/to/my-key", "my-bucket", "path/to/my-key", false},
		{"gs://my-bucket/", "", "", true},
		{"gs://", "", "", true},
		{"http://my-bucket/my-key", "", "", true},
		{"gs://my-bucket", "", "", true},
	}

	for _, tt := range tests {
		bucket, key, err := parseGCSURL(tt.url)
		if tt.err {
			if err == nil {
				t.Errorf("expected error for %s, got nil", tt.url)
			}
		} else {
			if err != nil {
				t.Errorf("expected no error for %s, got %v", tt.url, err)
			}
			if bucket != tt.bucket {
				t.Errorf("expected bucket %s, got %s", tt.bucket, bucket)
			}
			if key != tt.key {
				t.Errorf("expected key %s, got %s", tt.key, key)
			}
		}
	}
}

func TestS3DownloaderErrors(t *testing.T) {
	ctx := context.Background()
	d := NewS3Downloader("")

	assert.Equal(t, SourceTypeS3, d.Type())

	t.Run("Invalid URL", func(t *testing.T) {
		err := d.Download(ctx, "invalid-url", "dest")
		assert.Error(t, err)
	})

	t.Run("Failed to create directory", func(t *testing.T) {
		tmpFile, err := os.CreateTemp("", "blocked")
		require.NoError(t, err)
		defer os.Remove(tmpFile.Name())

		err = d.Download(ctx, "s3://bucket/key", filepath.Join(tmpFile.Name(), "subdir", "file"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create destination directory")
	})
}

func TestGCSDownloaderErrors(t *testing.T) {
	ctx := context.Background()
	d := NewGCSDownloader()

	assert.Equal(t, SourceTypeGCS, d.Type())

	t.Run("Invalid URL", func(t *testing.T) {
		err := d.Download(ctx, "invalid-url", "dest")
		assert.Error(t, err)
	})

	t.Run("Failed to create directory", func(t *testing.T) {
		tmpFile, err := os.CreateTemp("", "blocked-gcs")
		require.NoError(t, err)
		defer os.Remove(tmpFile.Name())

		err = d.Download(ctx, "gs://bucket/key", filepath.Join(tmpFile.Name(), "subdir", "file"))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "failed to create destination directory")
	})
}
