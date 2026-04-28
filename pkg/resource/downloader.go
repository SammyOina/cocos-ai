// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

// Package resource provides abstractions for downloading remote resources
// from various sources (OCI registries, S3, HTTP/HTTPS).
package resource

import (
	"context"
	"fmt"
	"sync"
)

// Downloader defines the interface for downloading resources from a remote source.
type Downloader interface {
	// Download fetches a resource from the given URL and writes it to destPath.
	// For OCI images, destPath is a directory. For S3/HTTP, destPath is a file path.
	Download(ctx context.Context, url string, destPath string) error

	// Type returns the source type identifier (e.g., "oci-image", "s3", "https", "http").
	Type() string
}

// Registry maps source type strings to Downloader implementations.
type Registry struct {
	mu          sync.RWMutex
	downloaders map[string]Downloader
}

// NewRegistry creates a new empty downloader registry.
func NewRegistry() *Registry {
	return &Registry{
		downloaders: make(map[string]Downloader),
	}
}

// Register adds a downloader to the registry for its declared type.
func (r *Registry) Register(d Downloader) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.downloaders[d.Type()] = d
}

// Get retrieves a downloader for the given source type.
func (r *Registry) Get(sourceType string) (Downloader, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	d, ok := r.downloaders[sourceType]
	if !ok {
		return nil, fmt.Errorf("unsupported source type: %s", sourceType)
	}
	return d, nil
}

// SupportedTypes returns a list of all registered source types.
func (r *Registry) SupportedTypes() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	types := make([]string, 0, len(r.downloaders))
	for t := range r.downloaders {
		types = append(types, t)
	}
	return types
}
