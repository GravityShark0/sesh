package dir

import "path/filepath"

func AlternatePath(s string) (altPath string) {
	if filepath.IsAbs(s) {
		return s
	}

	fullPath := FullPath(s)
	if fullPath != s {
		altPath = fullPath
	} else {
		if a, err := filepath.Abs(s); err == nil {
			altPath = a
		}
	}

	return altPath
}
