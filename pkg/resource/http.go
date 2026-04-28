// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

package resource

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

const (
	// SourceTypeHTTPS represents an HTTPS resource source.
	SourceTypeHTTPS = "https"
	// SourceTypeHTTP represents an HTTP resource source.
	SourceTypeHTTP = "http"

	httpTimeout = 5 * time.Minute
)

// HTTPDownloader downloads resources via HTTP/HTTPS.
type HTTPDownloader struct {
	client    *http.Client
	sourceTyp string
}

// NewHTTPSDownloader creates a new HTTPS downloader.
func NewHTTPSDownloader() *HTTPDownloader {
	return &HTTPDownloader{
		client: &http.Client{
			Timeout: httpTimeout,
		},
		sourceTyp: SourceTypeHTTPS,
	}
}

// NewHTTPDownloader creates a new HTTP downloader (insecure, for testing).
func NewHTTPDownloader() *HTTPDownloader {
	return &HTTPDownloader{
		client: &http.Client{
			Timeout: httpTimeout,
		},
		sourceTyp: SourceTypeHTTP,
	}
}

// Download fetches a resource from an HTTP/HTTPS URL and writes it to destPath.
func (h *HTTPDownloader) Download(ctx context.Context, url string, destPath string) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %w", err)
	}

	resp, err := h.client.Do(req)
	if err != nil {
		return fmt.Errorf("HTTP request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("HTTP request returned status %d: %s", resp.StatusCode, resp.Status)
	}

	// Ensure parent directory exists.
	if err := os.MkdirAll(filepath.Dir(destPath), 0o755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	f, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer f.Close()

	if _, err := io.Copy(f, resp.Body); err != nil {
		return fmt.Errorf("failed to write response body: %w", err)
	}

	return nil
}

// Type returns the source type identifier.
func (h *HTTPDownloader) Type() string {
	return h.sourceTyp
}
