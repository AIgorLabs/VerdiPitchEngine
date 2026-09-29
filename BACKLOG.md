---
agent_role: Core-Context
agent_weight: 4
asset_scope: Global
body_hash: "e2f65ae4434b71b8"
created: "2026-05-09"
dependencies: []
dev_stage: development
platform: CLI
priority: high
project_name: VerdiPitchEngine
status: active
tags:
    - dev-asset
    - docs
tech_stack:
    - Go
    - Shell
    - ffmpeg
updated: "2026-09-29"
version: 0.1.0
doc_id: "eb5fd50b-518a-4e2c-8290-6b3583a3388b"
frontmatter_hash: "5eb53390141eeb16"
---

# VerdiPitchEngine Backlog

Master continuous backlog tracking for **VerdiPitchEngine** (High-Fidelity Audio Pitch-Shifting Service).

---

## 🧊 Icebox & Future Operations

### [VPE-20004] NAS I/O Optimization ([#4](https://github.com/AIgorLabs/VerdiPitchEngine/issues/4))
- **Priority**: Medium
- **Type**: Feature
- **Origin**: Post-Mortem
- **Parent**: NONE
- **Blockers**: NONE
- **Dev Stage**: Planning
- **Status**: Pending
- **Manifesto Domain**: Latency & Performance
- **Task Master**: [.agent/tasks/VPE-20004-nas-i-o-optimization.md](.agent/tasks/VPE-20004-nas-i-o-optimization.md)
- **Description**: Optimize network-attached storage file throughput and I/O buffer allocation during batch FLAC conversion.

### [VPE-20005] Containerized E2E Runner ([#9](https://github.com/AIgorLabs/VerdiPitchEngine/issues/9))
- **Priority**: Medium
- **Type**: Feature
- **Origin**: Post-Mortem
- **Parent**: NONE
- **Blockers**: NONE
- **Dev Stage**: Planning
- **Status**: Pending
- **Manifesto Domain**: Quality Engineering
- **Task Master**: [.agent/tasks/VPE-20005-containerized-e2e-runner.md](.agent/tasks/VPE-20005-containerized-e2e-runner.md)
- **Description**: Implement automated containerized end-to-end test execution environment with pre-provisioned rubberband filter support.

### [VPE-20006] Spectral Analysis Validation ([#10](https://github.com/AIgorLabs/VerdiPitchEngine/issues/10))
- **Priority**: Medium
- **Type**: Feature
- **Origin**: Post-Mortem
- **Parent**: NONE
- **Blockers**: NONE
- **Dev Stage**: Planning
- **Status**: Pending
- **Manifesto Domain**: Quality Engineering
- **Task Master**: [.agent/tasks/VPE-20006-spectral-analysis-validation.md](.agent/tasks/VPE-20006-spectral-analysis-validation.md)
- **Description**: Add automated spectral analysis verification to validate 432 Hz frequency shifting accuracy without acoustic artifacts.

### [VPE-20009] Post-Processing Fidelity Validation ([#15](https://github.com/AIgorLabs/VerdiPitchEngine/issues/15))
- **Priority**: Medium
- **Type**: Feature
- **Origin**: Post-Mortem
- **Parent**: NONE
- **Blockers**: NONE
- **Dev Stage**: Planning
- **Status**: Pending
- **Manifesto Domain**: Quality Engineering
- **Task Master**: [.agent/tasks/VPE-20009-post-processing-fidelity-validation.md](.agent/tasks/VPE-20009-post-processing-fidelity-validation.md)
- **Description**: Establish automated bit-depth, dynamic range, and frequency spectrum verification for processed 24-bit audio tracks.

### [VPE-20011] Dynamic Context Timeout Bounding ([#21](https://github.com/AIgorLabs/VerdiPitchEngine/issues/21))
- **Priority**: Medium
- **Type**: Feature
- **Origin**: Post-Mortem
- **Parent**: NONE
- **Blockers**: NONE
- **Dev Stage**: Planning
- **Status**: Pending
- **Manifesto Domain**: Latency & Performance
- **Task Master**: [.agent/tasks/VPE-20011-dynamic-context-timeout-bounding.md](.agent/tasks/VPE-20011-dynamic-context-timeout-bounding.md)
- **Description**: Implement dynamic per-file context timeouts scaled by audio track duration to prevent conversion deadlocks.

### [VPE-20012] Headless / Daemon Log Redirection ([#22](https://github.com/AIgorLabs/VerdiPitchEngine/issues/22))
- **Priority**: Medium
- **Type**: Feature
- **Origin**: Post-Mortem
- **Parent**: NONE
- **Blockers**: NONE
- **Dev Stage**: Planning
- **Status**: Pending
- **Manifesto Domain**: Infrastructure & Architecture
- **Task Master**: [.agent/tasks/VPE-20012-headless-daemon-log-redirection.md](.agent/tasks/VPE-20012-headless-daemon-log-redirection.md)
- **Description**: Enable seamless switching between interactive MPB terminal output and structured headless daemon logging.

### [VPE-20014] Revert Operation (verdi-revert) ([#25](https://github.com/AIgorLabs/VerdiPitchEngine/issues/25))
- **Priority**: Medium
- **Type**: Feature
- **Origin**: Post-Mortem
- **Parent**: NONE
- **Blockers**: NONE
- **Dev Stage**: Planning
- **Status**: Pending
- **Manifesto Domain**: Quality Engineering
- **Task Master**: [.agent/tasks/VPE-20014-revert-operation-verdi-revert.md](.agent/tasks/VPE-20014-revert-operation-verdi-revert.md)
- **Description**: Build a safe rollback utility restoring original FLAC master files from hidden .verdi-backup directories.

### [VPE-20015] Prune Hidden Backups (verdi-prune) ([#26](https://github.com/AIgorLabs/VerdiPitchEngine/issues/26))
- **Priority**: Medium
- **Type**: Feature
- **Origin**: Post-Mortem
- **Parent**: NONE
- **Blockers**: NONE
- **Dev Stage**: Planning
- **Status**: Pending
- **Manifesto Domain**: Cost & Efficiency
- **Task Master**: [.agent/tasks/VPE-20015-prune-hidden-backups-verdi-prune.md](.agent/tasks/VPE-20015-prune-hidden-backups-verdi-prune.md)
- **Description**: Implement storage pruning utility to safely purge confirmed hidden backups and reclaim disk space.

### [VPE-20017] Dynamic Process Whitelisting Configuration ([#32](https://github.com/AIgorLabs/VerdiPitchEngine/issues/32))
- **Priority**: Medium
- **Type**: Feature
- **Origin**: Post-Mortem
- **Parent**: NONE
- **Blockers**: NONE
- **Dev Stage**: Planning
- **Status**: Pending
- **Manifesto Domain**: Security & Infrastructure
- **Task Master**: [.agent/tasks/VPE-20017-dynamic-process-whitelisting-configuration.md](.agent/tasks/VPE-20017-dynamic-process-whitelisting-configuration.md)
- **Description**: Support configurable process inspection lists via YAML configuration instead of hardcoded Roon/Plex process names.

### [VPE-20018] Containerized Log Rotation & Pruning ([#33](https://github.com/AIgorLabs/VerdiPitchEngine/issues/33))
- **Priority**: Medium
- **Type**: Feature
- **Origin**: Post-Mortem
- **Parent**: NONE
- **Blockers**: NONE
- **Dev Stage**: Planning
- **Status**: Pending
- **Manifesto Domain**: Infrastructure & Architecture
- **Task Master**: [.agent/tasks/VPE-20018-containerized-log-rotation-pruning.md](.agent/tasks/VPE-20018-containerized-log-rotation-pruning.md)
- **Description**: Implement automatic log rotation and retention policies for execution logs generated on Container Station.

### [VPE-20021] Cover Art Preservation ([#46](https://github.com/AIgorLabs/VerdiPitchEngine/issues/46))
- **Priority**: Medium
- **Type**: Feature
- **Origin**: Brainstorm
- **Parent**: NONE
- **Blockers**: NONE
- **Dev Stage**: Planning
- **Status**: Pending
- **Manifesto Domain**: Quality Engineering
- **Task Master**: [.agent/tasks/VPE-20021-cover-art-preservation.md](.agent/tasks/VPE-20021-cover-art-preservation.md)
- **Description**: Verify and guarantee byte-identical extraction and re-injection of embedded album art blocks across FLAC files.

### [VPE-20022] Go-Native FLAC Parsing Migration ([#35](https://github.com/AIgorLabs/VerdiPitchEngine/issues/35))
- **Priority**: Medium
- **Type**: Feature
- **Origin**: Post-Mortem
- **Parent**: NONE
- **Blockers**: NONE
- **Dev Stage**: Planning
- **Status**: Pending
- **Manifesto Domain**: Latency & Performance
- **Task Master**: [.agent/tasks/VPE-20022-go-native-flac-parsing-migration.md](.agent/tasks/VPE-20022-go-native-flac-parsing-migration.md)
- **Description**: Replace external metaflac and ffprobe binary invocations with a high-throughput native Go FLAC metadata parser.

### [VPE-20023] Zero-Copy Syscall Optimizations ([#36](https://github.com/AIgorLabs/VerdiPitchEngine/issues/36))
- **Priority**: Medium
- **Type**: Feature
- **Origin**: Post-Mortem
- **Parent**: NONE
- **Blockers**: NONE
- **Dev Stage**: Planning
- **Status**: Pending
- **Manifesto Domain**: Latency & Performance
- **Task Master**: [.agent/tasks/VPE-20023-zero-copy-syscall-optimizations.md](.agent/tasks/VPE-20023-zero-copy-syscall-optimizations.md)
- **Description**: Optimize inter-process piping between Go walker, conversion worker pool, and audio subprocesses using zero-copy buffers.

### [VPE-20024] Strategy Extensibility via YAML Configuration ([#38](https://github.com/AIgorLabs/VerdiPitchEngine/issues/38))
- **Priority**: Medium
- **Type**: Feature
- **Origin**: Post-Mortem
- **Parent**: NONE
- **Blockers**: NONE
- **Dev Stage**: Planning
- **Status**: Pending
- **Manifesto Domain**: Architecture & Tooling
- **Task Master**: [.agent/tasks/VPE-20024-strategy-extensibility-via-yaml-configuration.md](.agent/tasks/VPE-20024-strategy-extensibility-via-yaml-configuration.md)
- **Description**: Define dynamic pitch-shifting algorithms and filter graphs in declarative YAML configuration files.

### [VPE-20025] Metadata Byte-Level Parity Audit ([#39](https://github.com/AIgorLabs/VerdiPitchEngine/issues/39))
- **Priority**: Medium
- **Type**: Feature
- **Origin**: Post-Mortem
- **Parent**: NONE
- **Blockers**: NONE
- **Dev Stage**: Planning
- **Status**: Pending
- **Manifesto Domain**: Quality Engineering
- **Task Master**: [.agent/tasks/VPE-20025-metadata-byte-level-parity-audit.md](.agent/tasks/VPE-20025-metadata-byte-level-parity-audit.md)
- **Description**: Implement exhaustive byte-level diffing between origin and converted FLAC metadata blocks (excluding audio stream).

### [VPE-20026] Distributed Batch Orchestration ([#40](https://github.com/AIgorLabs/VerdiPitchEngine/issues/40))
- **Priority**: Medium
- **Type**: Feature
- **Origin**: Post-Mortem
- **Parent**: NONE
- **Blockers**: NONE
- **Dev Stage**: Planning
- **Status**: Pending
- **Manifesto Domain**: Infrastructure & Architecture
- **Task Master**: [.agent/tasks/VPE-20026-distributed-batch-orchestration.md](.agent/tasks/VPE-20026-distributed-batch-orchestration.md)
- **Description**: Enable multi-node worker distribution across local Mac workstation and remote QNAP Container Station instances.

### [VPE-20027] Migration to Google GenAI Go SDK ([#42](https://github.com/AIgorLabs/VerdiPitchEngine/issues/42))
- **Priority**: Medium
- **Type**: Feature
- **Origin**: Post-Mortem
- **Parent**: NONE
- **Blockers**: NONE
- **Dev Stage**: Planning
- **Status**: Pending
- **Manifesto Domain**: Architecture & Tooling
- **Task Master**: [.agent/tasks/VPE-20027-migration-to-google-genai-go-sdk.md](.agent/tasks/VPE-20027-migration-to-google-genai-go-sdk.md)
- **Description**: Migrate deprecated cloud.google.com/go/vertexai package in pkg/ai to the official google.golang.org/genai Go SDK.

### [VPE-20028] Dynamic Telemetry Credentials Configuration ([#43](https://github.com/AIgorLabs/VerdiPitchEngine/issues/43))
- **Priority**: Medium
- **Type**: Feature
- **Origin**: Post-Mortem
- **Parent**: NONE
- **Blockers**: NONE
- **Dev Stage**: Planning
- **Status**: Pending
- **Manifesto Domain**: Security & Infrastructure
- **Task Master**: [.agent/tasks/VPE-20028-dynamic-telemetry-credentials-configuration.md](.agent/tasks/VPE-20028-dynamic-telemetry-credentials-configuration.md)
- **Description**: Support dynamic GCP service account loading and environment variable fallbacks for telemetry authentication.

### [VPE-20029] Go Vulncheck Dependency Upgrades ([#44](https://github.com/AIgorLabs/VerdiPitchEngine/issues/44))
- **Priority**: Medium
- **Type**: Task
- **Origin**: Post-Mortem
- **Parent**: NONE
- **Blockers**: NONE
- **Dev Stage**: Planning
- **Status**: Pending
- **Manifesto Domain**: Security & Infrastructure
- **Task Master**: [.agent/tasks/VPE-20029-go-vulncheck-dependency-upgrades.md](.agent/tasks/VPE-20029-go-vulncheck-dependency-upgrades.md)
- **Description**: Remediate indirect dependency vulnerabilities surfaced by govulncheck through targeted go get module bumps.
