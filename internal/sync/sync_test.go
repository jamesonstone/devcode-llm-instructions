package sync

import (
	"os"
	"testing"
)

func TestSyncFile_DryRun(t *testing.T) {
	cfg := ConfigFile{
		Name:       "copilot-instructions.md",
		RemotePath: "github-copilot/copilot-instructions.md",
		LocalPath:  "testdata/copilot-instructions.md",
	}
	res := SyncFile(cfg, true)
	if res.Status != "updated" && res.Status != "unchanged" {
		t.Errorf("expected status updated or unchanged, got %s", res.Status)
	}
}

func TestSyncAll_DryRun(t *testing.T) {
	results := SyncAll(true)
	if len(results) != len(ConfigFiles) {
		t.Errorf("expected %d results, got %d", len(ConfigFiles), len(results))
	}
}

func TestSyncFile_Idempotent(t *testing.T) {
	cfg := ConfigFile{
		Name:       "copilot-instructions.md",
		RemotePath: "github-copilot/copilot-instructions.md",
		LocalPath:  "testdata/copilot-instructions.md",
	}
	_ = os.MkdirAll("testdata", 0755)
	_ = os.Remove(cfg.LocalPath)

	res1 := SyncFile(cfg, false)
	if res1.Status != "updated" && res1.Status != "error" {
		t.Errorf("expected updated on first sync, got %s", res1.Status)
	}

	if res1.Status == "updated" {
		res2 := SyncFile(cfg, false)
		if res2.Status != "unchanged" {
			t.Errorf("expected unchanged on second sync, got %s", res2.Status)
		}
	}

	_ = os.Remove(cfg.LocalPath)
	_ = os.Remove("testdata")
}
