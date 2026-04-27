// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

package resource

import (
	"context"

	"github.com/ultravioletrs/cocos/pkg/oci"
)

const (
	// SourceTypeOCIImage represents an OCI image resource source.
	SourceTypeOCIImage = "oci-image"
)

// OCIClient defines the interface for OCI image operations.
type OCIClient interface {
	PullAndDecrypt(ctx context.Context, source oci.ResourceSource, destDir string) error
	ToDockerArchive(ctx context.Context, ociDir, destFile string) error
}

// OCIDownloader adapts OCIClient to the Downloader interface.
// For OCI images, destPath is a directory where the OCI layout is written.
type OCIDownloader struct {
	client OCIClient
}


// NewOCIDownloader creates a new OCI downloader wrapping an OCI client.
func NewOCIDownloader(client OCIClient) *OCIDownloader {
	return &OCIDownloader{
		client: client,
	}
}

// Download pulls an OCI image to the destination directory.
// Note: For OCI images, encryption/decryption is handled by Skopeo + CoCo Keyprovider
// transparently via ocicrypt, so this just does the pull.
func (o *OCIDownloader) Download(ctx context.Context, url string, destDir string) error {
	source := oci.ResourceSource{
		Type: oci.ResourceTypeOCIImage,
		URI:  url,
		// Encryption handled separately by the caller who sets up Skopeo env.
		Encrypted: false,
	}
	return o.client.PullAndDecrypt(ctx, source, destDir)
}

// Type returns the source type identifier.
func (o *OCIDownloader) Type() string {
	return SourceTypeOCIImage
}

// Client returns the underlying OCIClient for OCI-specific operations
// like ToDockerArchive that aren't part of the generic Downloader interface.
func (o *OCIDownloader) Client() OCIClient {
	return o.client
}
