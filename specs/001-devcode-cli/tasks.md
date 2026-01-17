# Tasks: devcode CLI Tool

## Phase 1: Project Scaffold & Setup

- [x] T001: Initialize Go module and project structure (`go mod init`, create `cmd/devcode/`, `internal/sync/`)
- [x] T002: Add cobra dependency and scaffold root command (`devcode`)
- [x] T003: Implement `ConfigFile` and `SyncResult` structs in `internal/sync/model.go`
- [x] T004: Implement file download and sync logic in `internal/sync/sync.go` (reference: plan.md Data Model, CLI Contracts)
- [x] T005: Implement atomic file write and directory creation (reference: plan.md, see "Cross-platform directory creation")
- [x] T006: Implement idempotent sync (compare local/remote, only overwrite if changed)
- [x] T007: Implement error handling and reporting (reference: plan.md, "Error handling and reporting for CLI tools")

## Phase 2: CLI Commands & Flags

- [x] T008: Implement root command to sync all files (reference: plan.md CLI Contracts)
- [x] T009: Implement subcommands: `copilot`, `agents`, `warp`, `coderabbit` (reference: plan.md CLI Contracts)
- [x] T010: Implement global flags: `--dry-run`, `--version`, `--help` (reference: plan.md CLI Contracts)
- [x] T011: Implement sync status reporting ("Updated", "Up-to-date", "Error")

## Phase 3: Testing & Validation

- [x] T012: Write table-driven tests for sync logic (`internal/sync/`)
- [x] T013: Test CLI commands and flags (manual and automated)
- [x] T014: Validate cross-platform directory creation and file writing
- [x] T015: Validate idempotency (run sync twice, check no changes on second run)

## Phase 4: Migration & Documentation

- [x] T016: Remove `.envrc` and `example-direnv.sh` from repo (reference: plan.md Migration Steps)
- [x] T017: Update `README.md` to reference only devcode CLI (reference: plan.md Migration Steps)
- [x] T018: Update constitution to reflect CLI as sole distribution method (reference: plan.md Migration Steps)
- [x] T019: Remove all documentation references to direnv (reference: plan.md Migration Steps)

## Phase 5: Release & Distribution

- [x] T020: Add goreleaser config for cross-platform builds (optional, for binary releases)
- [x] T021: Tag release and verify `go install` works as expected

---

# Implementation Details Reference

- **Data Model, CLI Contracts, Directory Structure**: [plan.md](./plan.md)
- **Specification & Acceptance Criteria**: [spec.md](./spec.md)
- **Quality Checklist**: [checklists/requirements.md](./checklists/requirements.md)

# Notes

- Each task above should reference the relevant section in plan.md or spec.md for context.
- If implementation details are missing, update plan.md with clarifications as you go.
