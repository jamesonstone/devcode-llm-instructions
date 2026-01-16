#!/bin/sh
set -eu

# cleanup temp files on exit
trap 'rm -f "$COPILOT_FILE_TMP" "$AGENTS_FILE_TMP" "$WARP_FILE_TMP"' EXIT

dotenv_if_exists .env

# 1) copilot instructions (your current logic)
COPILOT_FILE_URL="https://raw.githubusercontent.com/jamesonstone/devcode-llm-instructions/refs/heads/main/github-copilot/copilot-instructions.md"
COPILOT_FILE_LOCAL="./.github/copilot-instructions.md"
COPILOT_FILE_TMP="/tmp/copilot-instructions.md"

mkdir -p "$(dirname "$COPILOT_FILE_LOCAL")"
curl -fsSL -o "$COPILOT_FILE_TMP" "$COPILOT_FILE_URL"

if [ -s "$COPILOT_FILE_TMP" ]; then
  if ! cmp -s "$COPILOT_FILE_TMP" "$COPILOT_FILE_LOCAL" 2>/dev/null; then
    cp "$COPILOT_FILE_TMP" "$COPILOT_FILE_LOCAL"
    echo "✅ Updated copilot-instructions.md"
  else
    echo "✅ copilot-instructions.md is up-to-date"
  fi
else
  echo "❌ Failed to download or empty file from $COPILOT_FILE_URL"
fi

# 2) spec-kit compatible rules: AGENTS.md + WARP.md
AGENTS_FILE_URL="https://raw.githubusercontent.com/jamesonstone/devcode-llm-instructions/refs/heads/main/spec-kit/AGENTS.md"
WARP_FILE_URL="https://raw.githubusercontent.com/jamesonstone/devcode-llm-instructions/refs/heads/main/spec-kit/WARP.md"

AGENTS_FILE_LOCAL="./AGENTS.md"
WARP_FILE_LOCAL="./WARP.md"

AGENTS_FILE_TMP="/tmp/AGENTS.md"
WARP_FILE_TMP="/tmp/WARP.md"

changed_any=0

curl -fsSL -o "$AGENTS_FILE_TMP" "$AGENTS_FILE_URL"
if [ -s "$AGENTS_FILE_TMP" ]; then
  if ! cmp -s "$AGENTS_FILE_TMP" "$AGENTS_FILE_LOCAL" 2>/dev/null; then
    cp "$AGENTS_FILE_TMP" "$AGENTS_FILE_LOCAL"
    echo "✅ Updated AGENTS.md"
    changed_any=1
  else
    echo "✅ AGENTS.md is up-to-date"
  fi
else
  echo "❌ Failed to download or empty file from $AGENTS_FILE_URL"
fi

curl -fsSL -o "$WARP_FILE_TMP" "$WARP_FILE_URL"
if [ -s "$WARP_FILE_TMP" ]; then
  if ! cmp -s "$WARP_FILE_TMP" "$WARP_FILE_LOCAL" 2>/dev/null; then
    cp "$WARP_FILE_TMP" "$WARP_FILE_LOCAL"
    echo "✅ Updated WARP.md"
    changed_any=1
  else
    echo "✅ WARP.md is up-to-date"
  fi
else
  echo "❌ Failed to download or empty file from $WARP_FILE_URL"
fi

if [ "$changed_any" -eq 1 ]; then
  echo "ℹ️ Warp: open this repo in Warp and run /init to (re)index + apply Project Rules."
fi
