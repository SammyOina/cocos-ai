// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

package gpu

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/ultravioletrs/cocos/pkg/attestation"
	"github.com/veraison/corim/corim"
)

const (
	DefaultVerifierBinary  = "nvidia-attestation-helper"
	defaultVerifierTimeout = 30 * time.Second
)

var _ attestation.Verifier = (*verifier)(nil)

type verifier struct {
	binaryPath         string
	timeout            time.Duration
	execCommandContext func(ctx context.Context, name string, arg ...string) *exec.Cmd
}

type evidenceEnvelope struct {
	Nonce string `json:"nonce"`
}

func NewVerifier(binaryPath string, timeout time.Duration) (attestation.Verifier, error) {
	if strings.TrimSpace(binaryPath) == "" {
		binaryPath = DefaultVerifierBinary
	}
	if timeout <= 0 {
		timeout = defaultVerifierTimeout
	}

	return &verifier{
		binaryPath:         binaryPath,
		timeout:            timeout,
		execCommandContext: exec.CommandContext,
	}, nil
}

func (v *verifier) VerifyWithCoRIM(report []byte, manifest *corim.UnsignedCorim) error {
	if len(report) == 0 {
		return fmt.Errorf("gpu evidence is empty")
	}

	nonceHex, err := evidenceNonce(report)
	if err != nil {
		return err
	}

	reqBody, err := json.Marshal(helperRequest{
		Mode:         "verify",
		NonceHex:     nonceHex,
		EvidenceJSON: append(json.RawMessage(nil), report...),
	})
	if err != nil {
		return fmt.Errorf("failed to marshal GPU verifier request: %w", err)
	}

	runCtx := context.Background()
	cancel := func() {}
	if v.timeout > 0 {
		runCtx, cancel = context.WithTimeout(runCtx, v.timeout)
	}
	defer cancel()

	stdout := &bytes.Buffer{}
	stderr := &bytes.Buffer{}

	cmd := v.execCommandContext(runCtx, v.binaryPath)
	cmd.Stdin = bytes.NewReader(reqBody)
	cmd.Stdout = stdout
	cmd.Stderr = stderr

	if err := cmd.Run(); err != nil {
		errMsg := strings.TrimSpace(stderr.String())
		if errMsg == "" {
			errMsg = err.Error()
		}
		return fmt.Errorf("gpu verifier helper failed: %s", errMsg)
	}

	var resp helperResponse
	if err := json.Unmarshal(stdout.Bytes(), &resp); err != nil {
		return fmt.Errorf("failed to decode GPU verifier response: %w", err)
	}
	if len(resp.ClaimsJSON) == 0 {
		return fmt.Errorf("gpu verifier response did not contain claims_json")
	}

	// NVIDIA attestation currently performs its own evidence-policy appraisal
	// and returns claims/detached EAT. We keep the attestation.Verifier
	// interface by treating manifest integration as a follow-up layer.
	_ = manifest

	return nil
}

func evidenceNonce(report []byte) (string, error) {
	var envelopes []evidenceEnvelope
	if err := json.Unmarshal(report, &envelopes); err != nil {
		return "", fmt.Errorf("failed to parse GPU evidence JSON: %w", err)
	}
	if len(envelopes) == 0 {
		return "", fmt.Errorf("gpu evidence did not contain any devices")
	}
	if strings.TrimSpace(envelopes[0].Nonce) == "" {
		return "", fmt.Errorf("gpu evidence nonce is missing")
	}

	return envelopes[0].Nonce, nil
}

// SetExecCommandContext allows tests to inject a mock exec.CommandContext.
func (v *verifier) SetExecCommandContext(cmdFunc func(ctx context.Context, name string, arg ...string) *exec.Cmd) {
	v.execCommandContext = cmdFunc
}
