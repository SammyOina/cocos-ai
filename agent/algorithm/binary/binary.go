// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0
package binary

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"

	"github.com/ultravioletrs/cocos/agent/algorithm"
	"github.com/ultravioletrs/cocos/agent/events"
	"github.com/ultravioletrs/cocos/pkg/socket"
)

const (
	socketPath = "unix_socket"
	pyRuntime  = "python3"
)

var _ algorithm.Algorithm = (*binary)(nil)

type binary struct {
	algoFile         string
	datasets         []string
	logger           *slog.Logger
	stderr           io.Writer
	stdout           io.Writer
	requirementsFile string
}

func New(logger *slog.Logger, eventsSvc events.Service, algoFile, requirementsFile string, datasets ...string) algorithm.Algorithm {
	return &binary{
		algoFile:         algoFile,
		datasets:         datasets,
		logger:           logger,
		stderr:           &algorithm.Stderr{Logger: logger, EventSvc: eventsSvc},
		stdout:           &algorithm.Stdout{Logger: logger},
		requirementsFile: requirementsFile,
	}
}

func (b *binary) Run() ([]byte, error) {
	var reqErr bytes.Buffer
	rcmd := exec.Command(pyRuntime, "-m", "pip", "install", "-r", b.requirementsFile)
	rcmd.Stderr = &reqErr
	if err := rcmd.Run(); err != nil {
		b.logger.Debug(reqErr.String())
		return nil, fmt.Errorf("error installing requirements: %v", err)
	}

	defer os.Remove(b.algoFile)
	defer func() {
		for _, file := range b.datasets {
			os.Remove(file)
		}
	}()
	listener, err := socket.StartUnixSocketServer(socketPath)
	if err != nil {
		return nil, fmt.Errorf("error creating stdout pipe: %v", err)
	}
	defer listener.Close()

	// Create channels for received data and errors
	dataChannel := make(chan []byte)
	errorChannel := make(chan error)

	var result []byte

	go socket.AcceptConnection(listener, dataChannel, errorChannel)

	args := append([]string{b.algoFile, socketPath}, b.datasets...)
	cmd := exec.Command(pyRuntime, args...)
	cmd.Stderr = b.stderr
	cmd.Stdout = b.stdout

	if err := cmd.Start(); err != nil {
		return nil, fmt.Errorf("error starting algorithm: %v", err)
	}

	if err := cmd.Wait(); err != nil {
		return nil, fmt.Errorf("algorithm execution error: %v", err)
	}

	select {
	case result = <-dataChannel:
		return result, nil
	case err = <-errorChannel:
		return nil, fmt.Errorf("error receiving data: %v", err)
	}
}
