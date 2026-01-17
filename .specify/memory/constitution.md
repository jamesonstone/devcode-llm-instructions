# <!--

# SYNC IMPACT REPORT

Version change: 0.0.0 → 1.0.0 (MAJOR - Initial ratification)

Modified principles: N/A (initial version)

Added sections:

- Core Principles (5 principles)
- Distribution Standards
- Development Workflow
- Governance

Removed sections: N/A (initial version)

Templates requiring updates:
✅ plan-template.md - No changes required
✅ spec-template.md - No changes required
✅ tasks-template.md - No changes required

# Follow-up TODOs: None

-->

# devcode-llm-instructions Constitution

## Core Principles

### I. Single Source of Truth

All LLM coding instructions MUST be centralized in this repository. Downstream projects sync from here—they do not maintain independent instruction sets. Local overrides are permitted only by disabling sync, and divergence MUST be intentional and documented.

**Rationale**: Centralization ensures consistency across all projects consuming these instructions. Without a single source, instruction drift creates unpredictable AI assistant behavior.

### II. Explicit Distribution

Configuration files MUST be distributed through explicit user action (CLI command, make target, or script invocation). Automatic background syncing (e.g., on every `cd`) is discouraged because it obscures when changes occur and complicates debugging.

**Rationale**: Developers should know when their tooling configuration changes. Implicit updates create confusion when AI behavior shifts unexpectedly.

### III. Idempotent Sync Operations

All sync operations MUST be idempotent and non-destructive. Running sync multiple times produces identical results. Files are only overwritten when content differs. Sync MUST report what changed (or confirm nothing changed) on every invocation.

**Rationale**: Safe operations encourage frequent syncing. Developers should never fear running the sync command.

### IV. Zero Runtime Dependencies

Distributed instruction files MUST work without requiring additional tooling beyond what the target AI assistant provides. No custom parsers, no runtime transformations, no compilation steps. The files are the final artifact.

**Rationale**: Complexity in the consumption path creates friction. Every project should be able to use these instructions immediately after sync.

### V. Transparent Change History

All changes to instruction files MUST be tracked in version control with meaningful commit messages. Breaking changes (changes that alter AI behavior significantly) MUST be documented in release notes.

**Rationale**: When AI behavior changes, developers need to trace the cause. Version control provides the audit trail.

## Distribution Standards

### Supported Configurations

| File                      | Purpose                     | Target Location                   |
| ------------------------- | --------------------------- | --------------------------------- |
| `copilot-instructions.md` | GitHub Copilot instructions | `.github/copilot-instructions.md` |
| `AGENTS.md`               | Spec-kit agent guidance     | `AGENTS.md`                       |
| `WARP.md`                 | Warp terminal rules         | `WARP.md`                         |
| `.coderabbit.yaml`        | CodeRabbit review config    | `.coderabbit.yaml`                |

### File Format Requirements

- Markdown files MUST use CommonMark format
- YAML files MUST include schema references where available
- All files MUST be UTF-8 encoded without BOM
- Line endings MUST be LF (not CRLF)

### Versioning Policy

Distribution follows semantic versioning:

- **MAJOR**: Breaking changes to instruction structure or removal of established patterns
- **MINOR**: New instruction files, new supported tools, expanded guidance
- **PATCH**: Typo fixes, clarifications, non-behavioral changes

## Development Workflow

### Adding New Instructions

1. Create instruction file in appropriate tool-specific directory (e.g., `github-copilot/`, `code-rabbit/`)
2. Update distribution mechanism (CLI tool or sync script) to include the new file
3. Document the file in this constitution's Distribution Standards table
4. Test sync in a clean project directory

### Modifying Existing Instructions

1. Make changes in the source file within this repository
2. Run sync in a test project to verify changes apply correctly
3. For behavioral changes, document in commit message what AI behavior is expected to change
4. Consider MAJOR version bump if change significantly alters AI assistant behavior

### Quality Gates

- All instruction files MUST be valid for their format (valid Markdown, valid YAML)
- Sync scripts MUST pass ShellCheck without errors
- Changes SHOULD be tested in at least one downstream project before merge

## Governance

This constitution supersedes all other documentation when conflicts arise. The constitution defines the rules; other documents (README, comments, etc.) provide guidance within those rules.

### Amendment Process

1. Propose change via pull request with rationale
2. Update constitution version according to semantic versioning rules
3. Update `LAST_AMENDED_DATE` to the merge date
4. Propagate changes to dependent templates if structure changes

### Compliance

- Pull requests adding new instruction files MUST update this constitution
- Breaking changes MUST increment MAJOR version
- All contributors are expected to read and follow this constitution

**Version**: 1.0.0 | **Ratified**: 2026-01-16 | **Last Amended**: 2026-01-16
