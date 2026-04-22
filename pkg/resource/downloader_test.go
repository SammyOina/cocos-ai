// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

package resource

import (
	"context"
	"testing"
)

type dummyDownloader struct {
	typ string
}

func (d *dummyDownloader) Download(ctx context.Context, url, destPath string) error {
	return nil
}

func (d *dummyDownloader) Type() string {
	return d.typ
}

func TestRegistry(t *testing.T) {
	reg := NewRegistry()

	// Initially empty
	if len(reg.SupportedTypes()) != 0 {
		t.Fatalf("expected 0 supported types, got %d", len(reg.SupportedTypes()))
	}

	// Register a downloader
	d1 := &dummyDownloader{typ: "test1"}
	reg.Register(d1)

	if len(reg.SupportedTypes()) != 1 {
		t.Fatalf("expected 1 supported type, got %d", len(reg.SupportedTypes()))
	}

	// Get the downloader
	got, err := reg.Get("test1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != d1 {
		t.Fatalf("expected to get identical downloader")
	}

	// Unknown type
	_, err = reg.Get("test2")
	if err == nil {
		t.Fatalf("expected error for unknown type")
	}
}
