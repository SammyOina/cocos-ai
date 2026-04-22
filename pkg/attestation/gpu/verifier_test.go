// Copyright (c) Ultraviolet
// SPDX-License-Identifier: Apache-2.0

package gpu

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/veraison/corim/corim"
)

func fakeVerifierExecCommandContext(_ context.Context, name string, arg ...string) *exec.Cmd {
	args := append([]string{"-test.run=TestGPUVerifierHelperProcess", "--", name}, arg...)
	cmd := exec.Command(os.Args[0], args...)
	cmd.Env = append(os.Environ(), "GO_WANT_GPU_VERIFIER_PROCESS=1")
	return cmd
}

func TestGPUVerifierHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_GPU_VERIFIER_PROCESS") != "1" {
		return
	}

	args := os.Args
	for i := range args {
		if args[i] == "--" {
			args = args[i+1:]
			break
		}
	}

	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "missing verifier binary name")
		os.Exit(2)
	}

	switch args[0] {
	case "verifier-error":
		fmt.Fprintln(os.Stderr, "simulated verifier failure")
		os.Exit(1)
	case "verifier-invalid-json":
		fmt.Fprintln(os.Stdout, "{not-json")
		os.Exit(0)
	case "verifier-empty-claims":
		fmt.Fprintln(os.Stdout, `{"detached_eat_json":{"overall_result":true}}`)
		os.Exit(0)
	default:
		var req helperRequest
		if err := json.NewDecoder(os.Stdin).Decode(&req); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		if req.Mode != "verify" {
			fmt.Fprintln(os.Stderr, "unexpected verifier mode")
			os.Exit(1)
		}
		if req.NonceHex == "" {
			fmt.Fprintln(os.Stderr, "nonce not propagated to verifier helper")
			os.Exit(1)
		}
		if !json.Valid(req.EvidenceJSON) {
			fmt.Fprintln(os.Stderr, "invalid evidence_json payload")
			os.Exit(1)
		}
		if !containsNonce(req.EvidenceJSON, req.NonceHex) {
			fmt.Fprintln(os.Stderr, "nonce not propagated to verifier")
			os.Exit(1)
		}

		resp := helperResponse{
			ClaimsJSON:      json.RawMessage(`[{"x-nvidia-device-type":"gpu"}]`),
			DetachedEATJSON: json.RawMessage(`{"overall_result":true}`),
		}
		if err := json.NewEncoder(os.Stdout).Encode(resp); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			os.Exit(1)
		}
		os.Exit(0)
	}
}

func TestEvidenceNonce(t *testing.T) {
	nonce, err := evidenceNonce([]byte(`[{"nonce":"aabbcc"}]`))
	assert.NoError(t, err)
	assert.Equal(t, "aabbcc", nonce)

	_, err = evidenceNonce([]byte(`[]`))
	assert.ErrorContains(t, err, "did not contain any devices")

	_, err = evidenceNonce([]byte(`[{}]`))
	assert.ErrorContains(t, err, "nonce is missing")
}

func containsNonce(report json.RawMessage, nonce string) bool {
	var envelopes []evidenceEnvelope
	if err := json.Unmarshal(report, &envelopes); err != nil {
		return false
	}
	if len(envelopes) == 0 {
		return false
	}

	return envelopes[0].Nonce == nonce
}

func TestVerifierVerifyWithCoRIM(t *testing.T) {
	v, err := NewVerifier("verifier-success", 0)
	require.NoError(t, err)

	cmdVerifier, ok := v.(*verifier)
	require.True(t, ok)
	cmdVerifier.SetExecCommandContext(fakeVerifierExecCommandContext)

	report := []byte(`[{"nonce":"aabbcc","evidence":"abc","certificate":"def"}]`)
	err = v.VerifyWithCoRIM(report, &corim.UnsignedCorim{})
	assert.NoError(t, err)
}

func TestVerifierVerifyWithCoRIMErrors(t *testing.T) {
	tests := []struct {
		name      string
		binary    string
		report    []byte
		wantError string
	}{
		{
			name:      "empty report",
			report:    nil,
			wantError: "gpu evidence is empty",
		},
		{
			name:      "invalid json",
			report:    []byte(`{`),
			wantError: "failed to parse GPU evidence JSON",
		},
		{
			name:      "helper failure",
			binary:    "verifier-error",
			report:    []byte(`[{"nonce":"aabbcc"}]`),
			wantError: "gpu verifier helper failed",
		},
		{
			name:      "invalid verifier response",
			binary:    "verifier-invalid-json",
			report:    []byte(`[{"nonce":"aabbcc"}]`),
			wantError: "failed to decode GPU verifier response",
		},
		{
			name:      "missing claims",
			binary:    "verifier-empty-claims",
			report:    []byte(`[{"nonce":"aabbcc"}]`),
			wantError: "gpu verifier response did not contain claims_json",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "empty report" || tt.name == "invalid json" {
				v, err := NewVerifier("verifier-success", 0)
				require.NoError(t, err)
				err = v.VerifyWithCoRIM(tt.report, nil)
				assert.ErrorContains(t, err, tt.wantError)
				return
			}

			v, err := NewVerifier(tt.binary, 0)
			require.NoError(t, err)

			cmdVerifier, ok := v.(*verifier)
			require.True(t, ok)
			cmdVerifier.SetExecCommandContext(fakeVerifierExecCommandContext)

			err = v.VerifyWithCoRIM(tt.report, nil)
			assert.ErrorContains(t, err, tt.wantError)
		})
	}
}
