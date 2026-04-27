// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

package resource

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/ultravioletrs/cocos/pkg/oci"
)

type MockOCIClient struct {
	mock.Mock
}

func (m *MockOCIClient) PullAndDecrypt(ctx context.Context, source oci.ResourceSource, destDir string) error {
	args := m.Called(ctx, source, destDir)
	return args.Error(0)
}

func (m *MockOCIClient) ToDockerArchive(ctx context.Context, ociDir, destFile string) error {
	args := m.Called(ctx, ociDir, destFile)
	return args.Error(0)
}

func TestOCIDownloader(t *testing.T) {
	mockClient := new(MockOCIClient)
	downloader := NewOCIDownloader(mockClient)

	ctx := context.Background()
	url := "docker://example.com/image:latest"
	destDir := "/tmp/oci"

	t.Run("Download", func(t *testing.T) {
		expectedSource := oci.ResourceSource{
			Type:      oci.ResourceTypeOCIImage,
			URI:       url,
			Encrypted: false,
		}
		mockClient.On("PullAndDecrypt", ctx, expectedSource, destDir).Return(nil).Once()

		err := downloader.Download(ctx, url, destDir)
		assert.NoError(t, err)
		mockClient.AssertExpectations(t)
	})

	t.Run("Type", func(t *testing.T) {
		assert.Equal(t, SourceTypeOCIImage, downloader.Type())
	})

	t.Run("Client", func(t *testing.T) {
		assert.Equal(t, mockClient, downloader.Client())
	})
}
