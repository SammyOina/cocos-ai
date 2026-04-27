// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

package resource

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

func TestHTTPDownloader(t *testing.T) {
	testContent := "test content"

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/test" {
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write([]byte(testContent))
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer ts.Close()

	d := NewHTTPDownloader()
	if d.Type() != SourceTypeHTTP {
		t.Fatalf("expected type %s, got %s", SourceTypeHTTP, d.Type())
	}

	tmpDir := t.TempDir()
	destPath := filepath.Join(tmpDir, "downloaded.txt")

	ctx := context.Background()

	// Test successful download
	err := d.Download(ctx, ts.URL+"/test", destPath)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	content, err := os.ReadFile(destPath)
	if err != nil {
		t.Fatalf("failed to read downloaded file: %v", err)
	}
	if string(content) != testContent {
		t.Fatalf("expected content %q, got %q", testContent, string(content))
	}

	// Test 404
	err = d.Download(ctx, ts.URL+"/notfound", filepath.Join(tmpDir, "notfound.txt"))
	if err == nil {
		t.Fatalf("expected error for 404 response")
	}
}
