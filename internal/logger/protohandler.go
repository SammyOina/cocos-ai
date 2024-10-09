// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0
package logger

import (
	"context"
	"io"
	"log/slog"
	"sync"
	"time"

	"github.com/ultravioletrs/cocos/pkg/manager"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const retryInterval = 5 * time.Second

var _ slog.Handler = (*handler)(nil)

type handler struct {
	opts           slog.HandlerOptions
	w              io.Writer
	cmpID          string
	cachedMessages [][]byte
	mutex          sync.Mutex
	stopRetry      chan struct{}
}

//go:generate mockery --name io.Writer --output ./mocks --filename io_writer.go

func NewProtoHandler(conn io.Writer, opts *slog.HandlerOptions, cmpID string) slog.Handler {
	if opts == nil {
		opts = &slog.HandlerOptions{}
	}
	h := &handler{
		opts:           *opts,
		w:              conn,
		cmpID:          cmpID,
		cachedMessages: make([][]byte, 0),
		stopRetry:      make(chan struct{}),
	}

	go h.periodicRetry()

	return h
}

func (h *handler) Enabled(_ context.Context, l slog.Level) bool {
	minLevel := slog.LevelInfo
	if h.opts.Level != nil {
		minLevel = h.opts.Level.Level()
	}
	return l >= minLevel
}

func (h *handler) Handle(_ context.Context, r slog.Record) error {
	message := r.Message
	timestamp := timestamppb.New(r.Time)
	level := r.Level.String()

	chunkSize := 500
	numChunks := (len(message) + chunkSize - 1) / chunkSize

	for i := 0; i < numChunks; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(message) {
			end = len(message)
		}

		chunk := message[start:end]

		agentLog := manager.ClientStreamMessage{
			Message: &manager.ClientStreamMessage_AgentLog{
				AgentLog: &manager.AgentLog{
					Timestamp:     timestamp,
					Message:       chunk,
					Level:         level,
					ComputationId: h.cmpID,
				},
			},
		}

		b, err := proto.Marshal(&agentLog)
		if err != nil {
			return err
		}

		h.mutex.Lock()
		_, err = h.w.Write(b)
		if err != nil {
			h.cachedMessages = append(h.cachedMessages, b)
		}
		h.mutex.Unlock()
	}

	return nil
}

func (h *handler) periodicRetry() {
	ticker := time.NewTicker(retryInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			h.retrySendCachedMessages()
		case <-h.stopRetry:
			return
		}
	}
}

func (h *handler) retrySendCachedMessages() {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	tmp := [][]byte{}
	for _, msg := range h.cachedMessages {
		if _, err := h.w.Write(msg); err != nil {
			tmp = append(tmp, msg)
		}
	}
	h.cachedMessages = tmp
}

func (h *handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	panic("unimplemented")
}

func (h *handler) WithGroup(name string) slog.Handler {
	panic("unimplemented")
}

func (h *handler) Close() error {
	close(h.stopRetry)
	return nil
}
