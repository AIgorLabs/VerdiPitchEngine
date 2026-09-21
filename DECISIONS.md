---
project_name: "VerdiPitchEngine"
title: "VerdiPitchEngine Architectural Decision Records (ADRs)"
description: "Authoritative technical design decisions, architectural trade-offs, and resolutions established in VerdiPitchEngine."
version: "1.0.0"
status: "Active"
priority: "Critical"
dev_stage: "Production"
agent_role: "Core-Context"
agent_weight: 5
asset_scope: "Global"
contains_pii: false
tenant_id: "AI.GOR-ENTERPRISE"
platform: "CLI"
tech_stack: ["Go", "Shell", "ffmpeg", "librubberband"]
dependencies: ["000-ECOSYSTEM", "001-AGENT", "007-SEVEN-GATE-LIFECYCLE", "020-GIT", "025-CHANGELOG", "027-TASK", "050-DOCS", "051-MEMORY", "055-GAP", "060-AIENG", "100-CORE", "1000-KEYS"]
doc_id: "8c1e4d3f-5b6a-4c2d-9e0f-1a2b3c4d5e6f"
created: "2026-05-09"
updated: "2026-09-21"
body_hash: "e503956f7c2dd8a9"
frontmatter_hash: "cfd3c49e84272900"
tags: [dev-asset, antigravity-context, adr, decisions, audio, dsp]
---

# Architectural Decision Records (ADRs)

This document records the authoritative technical design decisions, architectural trade-offs, and resolutions established in **VerdiPitchEngine** (`VPE`), the High-Fidelity Audio Pitch-Shifting Service for the **AI.gorLabs Software Ecosystem**.

---

## ADR-001: Selection of Time-Scale Modification via librubberband and Pure Asetrate

* **Status**: Accepted
* **Date**: 2026-05-09
* **Referenced Rules**: [Rule 000-ECOSYSTEM](.agent/rules/000-aigorlabs-ecosystem-taxonomy.md), [Rule 100-CORE](.agent/rules/100-core-architecture.md)
* **Context**: Pitch-shifting digital audio libraries from 440 Hz down to 432 Hz (-31.76 cents) via basic resampling introduces severe phase smearing, alters duration, and crushes high-frequency harmonic detail.
* **Decision**: Implement a dual-strategy architecture supporting both `librubberband` (duration-preserving Time-Scale Modification) and pure mathematical `asetrate` (phase-perfect preservation). Execute conversion pipelines via bounded containerized worker pools with zero-trust `:ro` input mounts.
* **Consequences**:
  - Delivers audiophile-grade high-fidelity transformations across 16-bit and 24-bit studio masters.
  - Eliminates real-time DSP CPU overhead on residential streaming endpoints (Roon Core, BluOS).
  - Preserves container isolation and host filesystem integrity.

---

## ADR-002: In-Place Audio Asset Modification and Filesystem Timestamp Mirroring

* **Status**: Accepted
* **Date**: 2026-05-15
* **Referenced Rules**: [Rule 100-CORE](.agent/rules/100-core-architecture.md)
* **Context**: Side-by-side library mirroring doubled storage requirements on NAS volumes and forced Roon and Plex Media Server to re-index tracks as duplicate albums, losing historical metadata, playlists, and user favorites.
* **Decision**: Adopt an in-place modification architecture where original 440 Hz FLAC assets are moved into hidden dot-prefixed backup directories (`.[440 Hz]`) within the existing album hierarchy. Mirror exact filesystem timestamps (`ModTime`, `AccessTime`) using `os.Chtimes` before final atomic moves.
* **Consequences**:
  - Prevents Roon from misidentifying converted files as newly added tracks.
  - Halves secondary storage consumption while maintaining full backup reversion paths.
  - Preserves album folder identity and database references seamlessly.

---

## ADR-003: Pure Math 24-Bit Lossless Audio Transformation and Metaflac Stream Parity

* **Status**: Accepted
* **Date**: 2026-05-20
* **Referenced Rules**: [Rule 060-AIENG](.agent/rules/060-ai-engineering-manifesto.md), [Rule 100-CORE](.agent/rules/100-core-architecture.md)
* **Context**: High-resolution studio masters (24-bit/96kHz, 192kHz) risk silent bit-depth truncation when standard FFmpeg default filters downsample audio to 16-bit PCM. Furthermore, standard transcoder tags strip MusicBrainz IDs and high-resolution album artwork blocks.
* **Decision**: Probe source media using `ffprobe` to strictly enforce bit-depth preservation into native `s32` format pipelines. Utilize `metaflac` byte-level block copying to ensure 1:1 metadata, lyrics, and picture block preservation.
* **Consequences**:
  - Guarantees zero loss in dynamic range or acoustic clarity.
  - Naturally strips proprietary MQA encoding layers while restoring true 24-bit linear PCM audio.
  - Preserves complete tag parity across modern digital audio players.

---

## ADR-004: 7-Gate Agentic Development Lifecycle & Sovereign Directives

* **Status**: Accepted
* **Date**: 2026-08-17
* **Referenced Rules**: [Rule 007-SEVEN-GATE-LIFECYCLE](.agent/rules/007-seven-gate-lifecycle.md), [Rule 001-AGENT](.agent/rules/001-agent-directives.md)
* **Context**: Autonomous AI agent operations in DSP and core conversion services require rigorous multi-stage verification to prevent audio corruption, resource leaks, or governance drift.
* **Decision**: Enforce the **7-Gate Agentic Development Lifecycle** for all VerdiPitchEngine changes, requiring systematic branch validation (Gate 1), implementation planning (Gate 2), atomic commits (Gate 3), quality gates via `mage check` (Gate 4), and post-mortem cognitive attribution (Gate 5-7).
* **Consequences**:
  - Guarantees 100% test coverage and build reproducibility.
  - Prevents regressions in concurrent audio pipelines.
  - Enforces deterministic, high-assurance agent operation.

---

## ADR-005: Trunk-Based Development, Rule 020-GIT Integrity & Zero-Touch CHANGELOG

* **Status**: Accepted
* **Date**: 2026-08-17
* **Referenced Rules**: [Rule 020-GIT](.agent/rules/020-git-standards.md), [Rule 025-CHANGELOG](.agent/rules/025-changelog-policy.md)
* **Context**: Manual edits to `CHANGELOG.md` cause merge conflicts across short-lived branches and violate cryptographic hash verification standards.
* **Decision**: Adopt trunk-based development with short-lived feature branches (`feat/VPE-*`, `fix/VPE-*`). Strictly prohibit manual edits or staging of `CHANGELOG.md`, delegating all changelog management exclusively to the automated `gen-release` pipeline.
* **Consequences**:
  - Eliminates changelog merge conflicts.
  - Preserves untampered release notes traceable to individual Git commits.
  - Upholds ecosystem-wide release engineering standards.

---

## ADR-006: Master Task Dossier Governance & Operational Gap Analysis

* **Status**: Accepted
* **Date**: 2026-08-18
* **Referenced Rules**: [Rule 027-TASK](.agent/rules/027-task-nomenclature.md), [Rule 055-GAP](.agent/rules/055-gap-analysis-output.md)
* **Context**: Audio engineering features and containerized NAS deployment tasks require precise specification and dependency tracking to avoid scope creep and broken DAG relationships.
* **Decision**: Adopt structured Master Task Dossiers (`.agent/tasks/VPE-*.md`) with formal provenance tracking (`Origin: Gap-Analysis | SWOT-Analysis | Root-Cause | Triage`) and continuous DAG cycle verification in `arc doctor`.
* **Consequences**:
  - Full end-to-end traceability for all DSP and toolchain enhancements.
  - Explicit blocker mapping prevents circular execution deadlocks.
  - Seamless synchronization with central ecosystem milestones.

---

## ADR-007: Dual-Stream Documentation Architecture & Semantic Memory Ingestion

* **Status**: Accepted
* **Date**: 2026-08-20
* **Referenced Rules**: [Rule 050-DOCS](.agent/rules/050-environmental-documentation.md), [Rule 051-MEMORY](.agent/rules/051-semantic-memory.md)
* **Context**: Operational runbooks for NAS deployment must remain separated from internal DSP architecture documentation, while historical resolutions need cross-agent semantic searchability.
* **Decision**: Partition documentation into operational runbooks (`RUNBOOK.md`) and technical guides, while publishing architectural milestones and post-mortems into OpenBrain (`pgvector`) semantic memory.
* **Consequences**:
  - Clear operational separation for NAS administrators vs DSP developers.
  - Autonomous agents leverage vector memory to recall past DSP optimizations.
  - Guarantees persistent knowledge retention across agent sessions.

---

## ADR-008: Dual SHA-256 Frontmatter & Body Cryptographic Integrity

* **Status**: Accepted
* **Date**: 2026-08-22
* **Referenced Rules**: [Rule 1000-KEYS](.agent/rules/1000-antigravity-keys.md)
* **Context**: Ecosystem security and governance mandates require all Markdown documentation, agent cards, and task records to be mathematically verifiable against tampering.
* **Decision**: Enforce Rule 1000-KEYS dual SHA-256 signatures (`body_hash` and `frontmatter_hash`) across all Markdown documents, verified automatically via `arc fixmarkdown` and pre-commit quality gates.
* **Consequences**:
  - Ensures absolute non-repudiation and tamper detection.
  - Automatically identifies untracked file modifications or schema drifts.
  - Fail-closes on malformed metadata during CI and local doctor checks.
