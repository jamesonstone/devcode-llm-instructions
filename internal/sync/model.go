package sync

// ConfigFile represents a file to sync from the central repo
// It contains the display name, remote path, and local target path.
type ConfigFile struct {
	Name      string // display name
	RemotePath string // path in source repo
	LocalPath  string // destination path in target project
}

// SyncResult represents the outcome of a sync operation
// It contains the file reference, status, and error message if applicable.
type SyncResult struct {
	File      ConfigFile
	Status    string // "updated", "unchanged", "error"
	ErrorMsg  string // non-empty if Status == "error"
}
