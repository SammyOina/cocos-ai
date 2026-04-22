use anyhow::Context;
use nv_attestation_sdk::{GpuEvidence, GpuEvidenceSource, Nonce, NvatSdk};
use serde::{Deserialize, Serialize};
use serde_json::Value;
use std::io::{stdin, stdout, Write};

#[derive(Debug, Deserialize)]
struct Request {
    nonce_hex: String,
}

#[derive(Debug, Serialize)]
struct Response {
    vendor: &'static str,
    evidence_format: &'static str,
    evidence_json: Value,
}

fn main() -> anyhow::Result<()> {
    let _sdk = NvatSdk::init_default().context("failed to initialize NVIDIA attestation SDK")?;

    let req: Request =
        serde_json::from_reader(stdin().lock()).context("failed to decode helper request")?;
    let nonce = Nonce::from_hex(&req.nonce_hex).context("failed to parse nonce")?;

    let evidence_source =
        GpuEvidenceSource::create_nvml().context("failed to create NVML evidence source")?;
    let evidence = GpuEvidence::collect(&evidence_source, Some(&nonce))
        .context("failed to collect evidence")?;
    let evidence_json = evidence
        .to_json()
        .context("failed to serialize GPU evidence to JSON")?;
    let evidence_json: Value =
        serde_json::from_str(&evidence_json).context("failed to parse serialized evidence JSON")?;

    let resp = Response {
        vendor: "nvidia",
        evidence_format: "nvat-json",
        evidence_json,
    };

    let mut out = stdout().lock();
    serde_json::to_writer(&mut out, &resp).context("failed to write helper response")?;
    let _ = out.write_all(b"\n");

    Ok(())
}
