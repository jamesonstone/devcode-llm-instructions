package sync

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

const repoBase = "https://raw.githubusercontent.com/jamesonstone/devcode-llm-instructions/main"

var ConfigFiles = []ConfigFile{
	{"copilot-instructions.md", "tool_templates/github-copilot/copilot-instructions.md", ".github/copilot-instructions.md"},
	{"AGENTS.md", "tool_templates/spec-kit/AGENTS.md", "AGENTS.md"},
	{"WARP.md", "tool_templates/spec-kit/WARP.md", "WARP.md"},
	{".coderabbit.yaml", "tool_templates/code-rabbit/.coderabbit.yaml", ".coderabbit.yaml"},
}

// SyncFile downloads and syncs a single config file. If dryRun is true, no files are written.
func SyncFile(cfg ConfigFile, dryRun bool) SyncResult {
	url := repoBase + "/" + cfg.RemotePath
	resp, err := http.Get(url)
	if err != nil {
		return SyncResult{cfg, "error", fmt.Sprintf("fetch failed: %v", err)}
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return SyncResult{cfg, "error", fmt.Sprintf("HTTP %d", resp.StatusCode)}
	}
	remoteData, err := io.ReadAll(resp.Body)
	if err != nil {
		return SyncResult{cfg, "error", fmt.Sprintf("read failed: %v", err)}
	}
	localData, err := os.ReadFile(cfg.LocalPath)
	if err == nil && sha256.Sum256(localData) == sha256.Sum256(remoteData) {
		return SyncResult{cfg, "unchanged", ""}
	}
	if dryRun {
		return SyncResult{cfg, "updated", "(dry run)"}
	}
	if dir := filepath.Dir(cfg.LocalPath); dir != "." {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return SyncResult{cfg, "error", fmt.Sprintf("mkdir failed: %v", err)}
		}
	}
	if err := os.WriteFile(cfg.LocalPath, remoteData, 0644); err != nil {
		return SyncResult{cfg, "error", fmt.Sprintf("write failed: %v", err)}
	}
	return SyncResult{cfg, "updated", ""}
}

// SyncAll syncs all config files. Returns a slice of SyncResult.
func SyncAll(dryRun bool) []SyncResult {
	results := make([]SyncResult, 0, len(ConfigFiles))
	for _, cfg := range ConfigFiles {
		results = append(results, SyncFile(cfg, dryRun))
	}
	return results
}
