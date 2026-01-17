package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const hookMarkerStart = "# >>> devcode auto-sync >>>"
const hookMarkerEnd = "# <<< devcode auto-sync <<<"

const zshHook = `# >>> devcode auto-sync >>>
# auto-sync devcode on directory change
devcode_auto_sync() {
  command -v devcode >/dev/null 2>&1 || return
  [[ -f .github/copilot-instructions.md || -f AGENTS.md || -f WARP.md || -f .coderabbit.yaml || -f .devcode ]] || return
  devcode >/dev/null 2>&1 &
}
autoload -Uz add-zsh-hook
add-zsh-hook chpwd devcode_auto_sync
devcode_auto_sync
# <<< devcode auto-sync <<<
`

const bashHook = `# >>> devcode auto-sync >>>
# auto-sync devcode on directory change
devcode_auto_sync() {
  command -v devcode >/dev/null 2>&1 || return
  [[ -f .github/copilot-instructions.md || -f AGENTS.md || -f WARP.md || -f .coderabbit.yaml || -f .devcode ]] || return
  devcode >/dev/null 2>&1 &
}
cd() {
  builtin cd "$@" && devcode_auto_sync
}
devcode_auto_sync
# <<< devcode auto-sync <<<
`

var installHookCmd = &cobra.Command{
	Use:   "install-hook",
	Short: "Install shell hook for automatic sync on directory change",
	Long: `Install a shell hook that automatically runs devcode when you cd into a project.

The hook is added to ~/.zshrc (for zsh) or ~/.bashrc (for bash).
It only runs devcode in directories that contain devcode-managed files.

To uninstall, run: devcode install-hook --uninstall`,
	Run: func(cmd *cobra.Command, args []string) {
		uninstall, _ := cmd.Flags().GetBool("uninstall")
		shell := detectShell()

		if shell == "" {
			fmt.Println("❌ Could not detect shell (zsh or bash)")
			os.Exit(1)
		}

		rcFile := shellRCFile(shell)
		if rcFile == "" {
			fmt.Printf("❌ Could not determine RC file for %s\n", shell)
			os.Exit(1)
		}

		if uninstall {
			if err := removeHook(rcFile); err != nil {
				fmt.Printf("❌ Failed to remove hook: %v\n", err)
				os.Exit(1)
			}
			fmt.Printf("✅ Removed devcode hook from %s\n", rcFile)
			fmt.Println("   Restart your shell or run: source " + rcFile)
			return
		}

		if hookExists(rcFile) {
			fmt.Printf("✅ Hook already installed in %s\n", rcFile)
			return
		}

		hook := zshHook
		if shell == "bash" {
			hook = bashHook
		}

		if err := appendHook(rcFile, hook); err != nil {
			fmt.Printf("❌ Failed to install hook: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✅ Installed devcode hook in %s\n", rcFile)
		fmt.Println("   Restart your shell or run: source " + rcFile)
	},
}

func init() {
	installHookCmd.Flags().Bool("uninstall", false, "Remove the shell hook")
	rootCmd.AddCommand(installHookCmd)
}

func detectShell() string {
	shell := os.Getenv("SHELL")
	if strings.Contains(shell, "zsh") {
		return "zsh"
	}
	if strings.Contains(shell, "bash") {
		return "bash"
	}
	return ""
}

func shellRCFile(shell string) string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	switch shell {
	case "zsh":
		return filepath.Join(home, ".zshrc")
	case "bash":
		return filepath.Join(home, ".bashrc")
	}
	return ""
}

func hookExists(rcFile string) bool {
	file, err := os.Open(rcFile)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), hookMarkerStart) {
			return true
		}
	}
	return false
}

func appendHook(rcFile, hook string) error {
	f, err := os.OpenFile(rcFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString("\n" + hook)
	return err
}

func removeHook(rcFile string) error {
	content, err := os.ReadFile(rcFile)
	if err != nil {
		return err
	}

	lines := strings.Split(string(content), "\n")
	var newLines []string
	inHook := false

	for _, line := range lines {
		if strings.Contains(line, hookMarkerStart) {
			inHook = true
			continue
		}
		if strings.Contains(line, hookMarkerEnd) {
			inHook = false
			continue
		}
		if !inHook {
			newLines = append(newLines, line)
		}
	}

	// remove trailing empty lines that might be left over
	for len(newLines) > 0 && newLines[len(newLines)-1] == "" {
		newLines = newLines[:len(newLines)-1]
	}

	return os.WriteFile(rcFile, []byte(strings.Join(newLines, "\n")+"\n"), 0644)
}
