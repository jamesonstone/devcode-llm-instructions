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
