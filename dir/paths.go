package dir

import (
	"os"
	"path/filepath"
	"strings"
)

// Turns a path into its absolute path
func AlternatePath(s string) string {
	switch {
	case s == "~/" || s == "~":
		homeDir, _ := os.UserHomeDir()
		return homeDir
	case filepath.IsAbs(s):
		return s
	case strings.HasPrefix(s, "~/"):
		if homeDir, err := os.UserHomeDir(); err == nil {
			return filepath.Join(homeDir, strings.TrimPrefix(s, "~/"))
		}
	default:
		if a, err := filepath.Abs(s); err == nil {
			return a
		}
	}
	return ""
}
