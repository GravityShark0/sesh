package dir

import (
	"os"
	"path/filepath"
	"strings"
)

func AlternatePath(s string) (altPath string) {
	if s == "~/" || s == "~" {
		homeDir, _ := os.UserHomeDir()
		return homeDir
	}

	if filepath.IsAbs(s) {
		return s
	}

	if strings.HasPrefix(s, "~/") {
		homeDir, err := os.UserHomeDir()
		if err == nil {
			return filepath.Join(homeDir, strings.TrimPrefix(s, "~/"))
		}
	}

	if a, err := filepath.Abs(s); err == nil {
		altPath = a
	}

	return altPath
}
