# NVIDIA Attestation Helper

This helper wraps NVIDIA's Rust attestation SDK low-level GPU evidence
collection flow and exposes a tiny JSON stdin/stdout protocol that the Go
attestation service can call.

## Request

The helper reads a single JSON object from stdin:

```json
{
  "nonce_hex": "aabbccdd"
}
```

## Response

On success it writes:

```json
{
  "vendor": "nvidia",
  "evidence_format": "nvat-json",
  "evidence_json": { "...": "..." }
}
```

`evidence_json` is the JSON emitted by `GpuEvidence::to_json()`.

## Build

Prerequisites:

- Rust 1.80+
- `libnvat.so.1`
- Clang/LLVM
- NVIDIA GPU driver with NVML support

If you are using a system-installed NVAT library:

```bash
export NVAT_USE_SYSTEM_LIB=1
cargo build --release
```

If you built NVAT locally, make sure the C library is installed or on
`LD_LIBRARY_PATH` before building or running the helper.

## Use With COCOS

Point the attestation service at the compiled binary:

```bash
export ATTESTATION_GPU_HELPER_PATH=/path/to/nvidia-attestation-helper
export ATTESTATION_GPU_HELPER_TIMEOUT=30s
```

When a helper path is configured, COCOS will attempt to collect GPU evidence
opportunistically. If the host does not expose a supported CC-capable NVIDIA
GPU, the attestation service skips GPU evidence and still returns the root
CPU/TEE attestation.
