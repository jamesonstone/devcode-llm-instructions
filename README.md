# 💻 devcode-llm-instructions

Centralized repository for LLM coding instructions with idempotent sync via the `devcode` CLI.

## Installation

```bash
go install github.com/jamesonstone/devcode@latest
```

## Usage

Sync all configuration files to your project:

```bash
devcode
```

Sync individual files:

```bash
devcode copilot     # .github/copilot-instructions.md
devcode agents      # AGENTS.md
devcode warp        # WARP.md
devcode coderabbit  # .coderabbit.yaml
```

Preview changes without writing files:

```bash
devcode --dry-run
```

## Auto-Sync (Optional)

Automatically sync when you `cd` into a project:

```bash
devcode install-hook
```

This adds a shell hook to your `~/.zshrc` or `~/.bashrc` that runs `devcode` in the background whenever you enter a directory containing devcode-managed files.

To remove the hook:

```bash
devcode install-hook --uninstall
```

## Features

- **Idempotent**: Running sync multiple times produces identical results
- **Non-destructive**: Files only overwritten when content differs
- **Status reporting**: Clear feedback on what changed (Updated, Up-to-date, Error)
- **Cross-platform**: Works on macOS, Linux, and Windows
- **Auto-sync**: Optional shell hook for automatic updates on directory change

## Supported Files

| File                        | Purpose                     | Target Location                   |
| --------------------------- | --------------------------- | --------------------------------- |
| `copilot-instructions.md`   | GitHub Copilot instructions | `.github/copilot-instructions.md` |
| `AGENTS.md`                 | Spec-kit agent guidance     | `AGENTS.md`                       |
| `WARP.md`                   | Warp terminal rules         | `WARP.md`                         |
| `.coderabbit.yaml`          | CodeRabbit review config    | `.coderabbit.yaml`                |

## Local Development

To customize instructions for a specific project, simply edit the synced files locally. Running `devcode` again will overwrite your changes—so only sync when you want to pull the latest centralized instructions.
