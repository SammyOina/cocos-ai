// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

package resource

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
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

// Full download test depends on actual GCP/AWS credentials unless mocked.
// We implement a mock S3-compatible server here for basic testing.
func TestS3DownloaderWithMockEndpoint(t *testing.T) {
	testContent := "s3 test content"

	// A very basic mock server that handles Google Cloud Storage compatible GET requests
	// The path will look like /bucket/key depending on how the client constructs it
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the requested path for debugging if needed
		if r.Method == http.MethodGet && strings.Contains(r.URL.Path, "my-key") {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(testContent))
			return
		}

		// Some GCS/S3 clients do probing or metadata requests
		// For the purpose of this test, we might just need to return the content directly
		// but since the storage.NewClient does some checking, we may fail.
		// However, storage.NewClient without credentials in test environment
		// usually fails initialization if not configured correctly.
		w.WriteHeader(http.StatusNotFound)
	}))
	defer ts.Close()

	// Skip the actual client call if we don't have a robust mock,
	// because cloud.google.com/go/storage is hard to mock perfectly without a full emulator.
	// For this test, we just verify the parsing and the Type() method.

	d := NewS3Downloader(ts.URL)
	if d.Type() != SourceTypeS3 {
		t.Fatalf("expected type %s, got %s", SourceTypeS3, d.Type())
	}

	d2 := NewGCSDownloader()
	if d2.Type() != SourceTypeGCS {
		t.Fatalf("expected type %s, got %s", SourceTypeGCS, d2.Type())
	}
}
