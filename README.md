# 💻 devcode-llm-instructions

Centralized Repository for Coding LLM Instructions Leveraging `direnv` for Change Detection

This project is built on top of `direnv`, a shell extension that allows you to load and unload environment variables based on the current directory. The aim of this project is to centralize coding instructions for new projects. To advance or iterate on project specific instructions, you can simply disable the `direnv` hook, make your changes, and then re-enable it.

For more information on `direnv`, visit [direnv.net](https://direnv.net).

🧪 This project is experimental 🚧

## How to Use

1. Download/Install `direnv`:
   - For macOS: `brew install direnv`
   - For Linux: Follow instructions on [direnv's website](https://direnv.net/docs/installation.html)
   - For Windows: Use WSL or follow [direnv's Windows guide](https://direnv.net/docs/installation.html#windows)
2. Create an empty `.envrc` file in your project directory:
   ```bash
   touch .envrc
   ```
3. If you're using copilot, you can copy the contents of the `example-direnv.sh` file from this repository into your `.envrc` file:
   ```bash
   cp [example-direnv.sh](https://raw.githubusercontent.com/jamesonstone/devcode-llm-instructions/main/example-direnv.sh) .envrc
   ```
4. Allow `direnv` to load the `.envrc` file:
   ```bash
   direnv allow .
   ```
5. Now, whenever you enter the directory, `direnv` will automatically load the environment variables defined in `.envrc` and check for new instructions in this repository.
6. If you want to make changes to the instructions, you can disable `direnv` from syncing these instructions by removing the `.envrc` file or commenting out the relevant lines in it. After making your changes, you can re-enable `direnv` by running:
   ```bash
   direnv allow .
   ```

## Example `.envrc` File

```bash
#!/usr/bin/sh

dotenv_if_exists .env # use the local .env file if it exists

COPILOT_FILE_URL="https://raw.githubusercontent.com/jamesonstone/devcode-llm-instructions/refs/heads/main/github-copilot/copilot-instructions.md"
COPILOT_FILE_LOCAL="./.github/copilot-instructions.md"
COPILOT_FILE_TMP="/tmp/copilot-instructions.md"

# Create the directory for COPILOT_FILE_LOCAL if it doesn't exist
COPILOT_FILE_LOCAL_DIR="$(dirname "$COPILOT_FILE_LOCAL")"
mkdir -p "$COPILOT_FILE_LOCAL_DIR"

curl -s -o "$COPILOT_FILE_TMP" "$COPILOT_FILE_URL"

if [ -s "$COPILOT_FILE_TMP" ]; then
  if ! cmp -s "$COPILOT_FILE_TMP" "$COPILOT_FILE_LOCAL"; then
    cp "$COPILOT_FILE_TMP" "$COPILOT_FILE_LOCAL"
    echo "✅ Updated copilot-instructions.md"
  else
    echo "✅ copilot-instructions.md is up-to-date"
  fi
else
  echo "❌ Failed to download or empty file from $COPILOT_FILE_URL"
fi
```
