---
project_name: "VerdiPitchEngine"
rule_id: "000-AGENTCARD"
title: "VerdiPitchEngine Universal AI Agent Operational Directive Card"
description: "Authoritative root operational contract and instructions for autonomous AI agents operating within the VerdiPitchEngine repository."
version: "1.0.0"
status: "Active"
priority: "Critical"
dev_stage: "Production"
agent_role: "Core-Context"
agent_weight: 5
asset_scope: "Global"
target_projects: ["VerdiPitchEngine"]
applies_to_paths: ["AGENTS.md"]
domains: ["Agent-Directive", "Audio", "DSP", "VerdiPitchEngine"]
contains_pii: false
tenant_id: "AI.GOR-ENTERPRISE"
platform: "CLI"
tech_stack: ["Go", "Shell", "ffmpeg", "librubberband"]
dependencies: ["000-ECOSYSTEM", "000-CONTEXT", "001-AGENT", "020-GIT", "021-BRANCH", "025-CHANGELOG", "027-TASK", "100-CORE", "900-GO", "1000-KEYS"]
gem_id: "VerdiCore"
gem_aliases: ["!Verdi", "!VerdiPitchEngine", "!VPE"]
gem_dependencies:
    - "01845d72-3999-4a34-a9b2-e0807a1650c2" # ArchitectNexus
    - "6fca3424-a350-4d91-ad9c-9f9d1e6bf897" # AI.gor-Core
openbrain_mounts:
    - "Atlas-Specs"
    - "OpenBrain-Memory"
    - "Lab-Gems-Personas"
code_authority:
    level: full
    allow_file_writes: true
    allow_tool_execution: true
    allow_pr_generation: true
    allow_commit_generation: true
    allowed_languages: ["go", "python", "markdown", "shell", "yaml", "json"]
    disallowed_paths:
        - "CHANGELOG.md"
        - ".git/**"
created: "2026-08-17"
updated: "2026-09-21"
tags: [agent-card, operational-contract, verdipitchengine, audio, dsp]
doc_id: "e4f5a6b7-c8d9-4e0f-1a2b-3c4d5e6f7a8b"
body_hash: "463091a3e84b5412"
frontmatter_hash: "80267b5e396a1ff7"
---

# VerdiPitchEngine Agent Operational Directive Card

You are operating inside **VerdiPitchEngine** (`VPE`), a **personal high-fidelity audio pitch-shifting service** managed by the **Archon Developer Toolchain & Central Architectural Governor**. Note that VerdiPitchEngine is an independent personal side-project (not part of the AI.gorLabs commercial services ecosystem), but strictly conforms to Archon governance, 1000-KEYS cryptographic standards, and developer workflows.

---

## 🛑 Canonical Invariants & Decoupling Boundaries ([Rule 000-ECOSYSTEM](.agent/rules/000-aigorlabs-ecosystem-taxonomy.md))

1. **Audio Processing Scope**:
   - **VerdiPitchEngine IS**: Containerized Go application for high-fidelity batch Time-Scale Modification (TSM) on lossless audio files (pitch-shifting from 440 Hz to 432 Hz).
   - **Zero-Trust Audio Safety**: Strictly read-only (`:ro`) source volume mount for master library files. Writes directed to isolated output volumes.

2. **Zero-Touch CHANGELOG ([Rule 025-CHANGELOG](.agent/rules/025-changelog-policy.md))**:
   - Never stage, edit, or format `CHANGELOG.md`.

3. **Cryptographic Signing ([Rule 1000-KEYS](.agent/rules/1000-antigravity-keys.md))**:
   - Recompute markdown hashes with `arc fixmarkdown <path>`.

4. **Quality Gate Verification**:
   - Run `mage check` or `go test -v -race ./...` before concluding tasks.

---

## 🛠️ Tooling Reference (`arc`)

- `arc doctor` — Validates environment health and symlink integrity.
- `arc init-task <ID>` — Scaffolds task implementation workspace.
- `arc fixmarkdown <path>` — Computes and synchronizes Rule 1000-KEYS hashes.
