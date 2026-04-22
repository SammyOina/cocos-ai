// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

package resource

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

const (
	// SourceTypeS3 represents an S3-compatible object storage source.
	SourceTypeS3 = "s3"
	// SourceTypeGCS represents a Google Cloud Storage source.
	SourceTypeGCS = "gcs"
)

// S3Downloader downloads resources from S3-compatible object storage.
// It uses the Google Cloud Storage client library with S3-compatible endpoints
// or can be configured for MinIO/AWS S3 via environment variables.
type S3Downloader struct {
	endpoint string // Optional custom endpoint for S3-compatible services (e.g., MinIO).
}

// NewS3Downloader creates a new S3 downloader.
// If endpoint is empty, standard AWS S3 environment credentials/config are used.
func NewS3Downloader(endpoint string) *S3Downloader {
	return &S3Downloader{
		endpoint: endpoint,
	}
}

// Download fetches a resource from an S3 URL (s3://bucket/key) and writes it to destPath.
func (s *S3Downloader) Download(ctx context.Context, url string, destPath string) error {
	bucket, key, err := parseS3URL(url)
	if err != nil {
		return err
	}

	// Ensure parent directory exists.
	if err := os.MkdirAll(filepath.Dir(destPath), 0o755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	// Use Google Cloud Storage client with S3-compatible XML API when endpoint is set.
	// For standard GCS, use default credentials.
	var opts []option.ClientOption
	if s.endpoint != "" {
		opts = append(opts, option.WithEndpoint(s.endpoint))
	}

	client, err := storage.NewClient(ctx, opts...)
	if err != nil {
		return fmt.Errorf("failed to create storage client: %w", err)
	}
	defer client.Close()

	reader, err := client.Bucket(bucket).Object(key).NewReader(ctx)
	if err != nil {
		return fmt.Errorf("failed to read object %s/%s: %w", bucket, key, err)
	}
	defer reader.Close()

	f, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer f.Close()

	if _, err := f.ReadFrom(reader); err != nil {
		return fmt.Errorf("failed to write object content: %w", err)
	}

	return nil
}

// Type returns the source type identifier.
func (s *S3Downloader) Type() string {
	return SourceTypeS3
}

// GCSDownloader downloads resources from Google Cloud Storage.
type GCSDownloader struct{}

// NewGCSDownloader creates a new GCS downloader.
func NewGCSDownloader() *GCSDownloader {
	return &GCSDownloader{}
}

// Download fetches a resource from a GCS URL (gs://bucket/key) and writes it to destPath.
func (g *GCSDownloader) Download(ctx context.Context, url string, destPath string) error {
	bucket, key, err := parseGCSURL(url)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(destPath), 0o755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("failed to create GCS client: %w", err)
	}
	defer client.Close()

	reader, err := client.Bucket(bucket).Object(key).NewReader(ctx)
	if err != nil {
		return fmt.Errorf("failed to read object gs://%s/%s: %w", bucket, key, err)
	}
	defer reader.Close()

	f, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer f.Close()

	if _, err := f.ReadFrom(reader); err != nil {
		return fmt.Errorf("failed to write object content: %w", err)
	}

	return nil
}

// Type returns the source type identifier.
func (g *GCSDownloader) Type() string {
	return SourceTypeGCS
}

// parseS3URL parses an S3 URL of the form s3://bucket/key.
func parseS3URL(url string) (bucket, key string, err error) {
	if !strings.HasPrefix(url, "s3://") {
		return "", "", fmt.Errorf("invalid S3 URL, expected s3://bucket/key, got: %s", url)
	}
	path := strings.TrimPrefix(url, "s3://")
	parts := strings.SplitN(path, "/", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("invalid S3 URL, expected s3://bucket/key, got: %s", url)
	}
	return parts[0], parts[1], nil
}

// parseGCSURL parses a GCS URL of the form gs://bucket/key.
func parseGCSURL(url string) (bucket, key string, err error) {
	if !strings.HasPrefix(url, "gs://") {
		return "", "", fmt.Errorf("invalid GCS URL, expected gs://bucket/key, got: %s", url)
	}
	path := strings.TrimPrefix(url, "gs://")
	parts := strings.SplitN(path, "/", 2)
	if len(parts) != 2 || parts[0] == "" || parts[1] == "" {
		return "", "", fmt.Errorf("invalid GCS URL, expected gs://bucket/key, got: %s", url)
	}
	return parts[0], parts[1], nil
}
