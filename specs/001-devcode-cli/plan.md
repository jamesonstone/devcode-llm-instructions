# Implementation Plan: devcode CLI Tool

**Feature Branch**: `001-devcode-cli`
**Created**: 2026-01-16
**Status**: Complete

## Technical Context

- **Language/Version**: Go 1.21+
- **CLI Framework**: cobra
- **Distribution**: `go install github.com/jamesonstone/devcode@latest`
- **Supported OS**: macOS, Linux, Windows
- **Dependencies**: cobra, stdlib only
- **Testing**: Go test, table-driven tests for sync logic
- **Target Platform**: Local developer machines
- **Project Type**: CLI tool
- **Performance Goals**: Sync all files in <5s, binary <10MB

## Constitution Check

- All distribution and migration requirements from spec.md are included
- No references to `direnv` or `.envrc` in code, docs, or workflow
- CLI is the only supported distribution mechanism

## Gates

- [x] Spec and checklist validated
- [x] Migration requirements present
- [x] No [NEEDS CLARIFICATION] markers

## Phase 0: Research

- [x] Best practices for cobra CLI structure
- [x] Go idioms for file download, comparison, and atomic write
- [x] Error handling and reporting for CLI tools
- [x] Cross-platform directory creation
- [x] Table-driven tests for CLI commands

## Phase 1: Design & Contracts

### Data Model

- **ConfigFile**: name, remotePath, localPath
- **SyncResult**: file, status (updated, unchanged, error), errorMsg

### CLI Contracts

- `devcode` (root): sync all files
- `devcode copilot`: sync only copilot-instructions.md
- `devcode agents`: sync only AGENTS.md
- `devcode warp`: sync only WARP.md
- `devcode coderabbit`: sync only .coderabbit.yaml
- Global flags: `--dry-run`, `--version`, `--help`

### Directory Structure

```text
cmd/devcode/
    main.go
    root.go
    copilot.go
    agents.go
    warp.go
    coderabbit.go
internal/sync/
    sync.go
    files.go
    model.go
    testdata/
        ...
```

## Quickstart

1. Install: `go install github.com/jamesonstone/devcode@latest`
2. Run: `devcode` (sync all files)
3. Run: `devcode copilot` (sync only copilot)
4. Run: `devcode --dry-run` (preview changes)

## Migration Steps

1. Remove `.envrc` and `example-direnv.sh` from repo
2. Update `README.md` to reference only devcode CLI
3. Update constitution to reflect CLI as sole distribution method
4. Remove all documentation references to direnv

## Validation

- All files sync as expected
- No direnv or .envrc references remain
- CLI passes all tests
- Documentation matches new workflow

---

Ready for implementation.
