# Feature Specification: devcode CLI Tool

**Feature Branch**: `001-devcode-cli`
**Created**: 2026-01-16
**Status**: Draft
**Input**: User description: "Go CLI tool to download LLM instruction files with cobra framework, default downloads all files, individual file options available"

## User Scenarios & Testing _(mandatory)_

### User Story 1 - Download All Configuration Files (Priority: P1)

As a developer setting up a new project, I want to run a single command to download all LLM instruction files so that my project is immediately configured with standardized AI assistant instructions.

**Why this priority**: This is the primary use case—most users want all files at once. Without this, the tool provides no value.

**Independent Test**: Run `devcode` in an empty directory and verify all four configuration files are created in their correct locations.

**Acceptance Scenarios**:

1. **Given** a directory without any devcode configuration files, **When** I run `devcode`, **Then** all four files are downloaded to their target locations (`.github/copilot-instructions.md`, `AGENTS.md`, `WARP.md`, `.coderabbit.yaml`)
2. **Given** a directory with existing devcode files that match the remote versions, **When** I run `devcode`, **Then** the tool reports all files are up-to-date without overwriting
3. **Given** a directory with outdated devcode files, **When** I run `devcode`, **Then** only the changed files are updated and the tool reports which files were updated

---

### User Story 2 - Download Individual Files (Priority: P2)

As a developer who only uses specific AI tools, I want to download individual configuration files so that I don't clutter my project with unused configurations.

**Why this priority**: Provides flexibility for users with specific needs. Builds on the core sync mechanism from P1.

**Independent Test**: Run `devcode copilot` and verify only `.github/copilot-instructions.md` is created/updated.

**Acceptance Scenarios**:

1. **Given** a directory without copilot instructions, **When** I run `devcode copilot`, **Then** only `.github/copilot-instructions.md` is created
2. **Given** a directory without agents configuration, **When** I run `devcode agents`, **Then** only `AGENTS.md` is created
3. **Given** a directory without warp configuration, **When** I run `devcode warp`, **Then** only `WARP.md` is created
4. **Given** a directory without coderabbit configuration, **When** I run `devcode coderabbit`, **Then** only `.coderabbit.yaml` is created

---

### User Story 3 - Global Installation via go install (Priority: P3)

As a developer, I want to install devcode globally using `go install` so that I can use it across all my projects without additional setup.

**Why this priority**: Distribution mechanism. Without it, the tool can't reach users, but the core functionality must work first.

**Independent Test**: Run `go install github.com/jamesonstone/devcode@latest` and verify `devcode` is available in PATH.

**Acceptance Scenarios**:

1. **Given** Go is installed and GOBIN is in PATH, **When** I run `go install github.com/jamesonstone/devcode@latest`, **Then** `devcode` command becomes available system-wide
2. **Given** devcode is installed, **When** I run `devcode --version`, **Then** the version number is displayed
3. **Given** devcode is installed, **When** I run `devcode --help`, **Then** usage information with all available subcommands is displayed

---

### User Story 4 - Dry Run Mode (Priority: P4)

As a developer, I want to preview what files would be changed before actually downloading them so that I can make informed decisions about updates.

**Why this priority**: Nice-to-have safety feature. Core functionality works without it.

**Independent Test**: Run `devcode --dry-run` and verify no files are created/modified but output shows what would happen.

**Acceptance Scenarios**:

1. **Given** a directory with outdated files, **When** I run `devcode --dry-run`, **Then** the tool reports which files would be updated without actually modifying them
2. **Given** a directory without any files, **When** I run `devcode copilot --dry-run`, **Then** the tool reports the file would be created without actually creating it

---

### Edge Cases

- What happens when the network is unavailable? Tool reports connection error clearly and exits with non-zero status.
- What happens when the user doesn't have write permissions to the target directory? Tool reports permission error with the specific path.
- What happens when GitHub raw content CDN is slow or times out? Tool uses a reasonable timeout (10 seconds) and reports timeout errors.
- What happens when the remote file is empty or returns an error? Tool skips that file, reports the error, and continues with other files.
- What happens when `.github/` directory doesn't exist? Tool creates it automatically before writing copilot-instructions.md.

## Requirements _(mandatory)_

### Functional Requirements

- **FR-001**: CLI MUST download files from `https://raw.githubusercontent.com/jamesonstone/devcode-llm-instructions/main/` when no branch is specified
- **FR-002**: CLI MUST support a root command (`devcode`) that downloads all configured files
- **FR-003**: CLI MUST support subcommands for individual files: `copilot`, `agents`, `warp`, `coderabbit`
- **FR-004**: CLI MUST compare local and remote file contents before overwriting (idempotent sync)
- **FR-005**: CLI MUST create parent directories if they don't exist (e.g., `.github/`)
- **FR-006**: CLI MUST report sync status for each file: "Updated", "Up-to-date", or "Error"
- **FR-007**: CLI MUST support `--version` flag showing semantic version
- **FR-008**: CLI MUST support `--help` flag with usage documentation
- **FR-009**: CLI MUST support `--dry-run` flag to preview changes without writing files
- **FR-010**: CLI MUST exit with code 0 on success, non-zero on any error
- **FR-011**: CLI MUST use cobra framework for command structure
- **FR-012**: CLI MUST be installable via `go install github.com/jamesonstone/devcode@latest`

### Migration Requirements (direnv → devcode)

- **MR-001**: Project MUST remove `.envrc` file from repository after CLI is functional
- **MR-002**: Project MUST remove `example-direnv.sh` file from repository after CLI is functional
- **MR-003**: Project MUST update `README.md` to replace all direnv references with devcode CLI instructions
- **MR-004**: Project MUST update constitution to reflect the new distribution mechanism
- **MR-005**: Documentation MUST NOT reference direnv as a distribution method after migration is complete

### Key Entities

- **ConfigFile**: Represents a file to sync—has name, remote path, local target path
- **SyncResult**: Outcome of a sync operation—file reference, status (updated/unchanged/error), error message if applicable

## Success Criteria _(mandatory)_

### Measurable Outcomes

- **SC-001**: Users can sync all files with a single command in under 5 seconds on standard broadband
- **SC-002**: Binary size remains under 10MB to ensure fast downloads
- **SC-003**: 100% of sync operations are idempotent—running twice produces identical filesystem state
- **SC-004**: Tool provides clear feedback—users can determine sync status within 2 seconds of command completion
- **SC-005**: Installation via `go install` completes in under 30 seconds
