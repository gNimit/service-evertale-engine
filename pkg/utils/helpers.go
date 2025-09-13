package utils

import (
	"os"
	"path/filepath"
)

// FindRootDir traverses up the directory tree to find the root directory containing the .env file.
func FindRootDir() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Navigate upwards until finding the directory containing the go.mod file
	for {
		if _, err := os.Stat(filepath.Join(cwd, ".env")); err == nil {
			return cwd, nil
		}
		parent := filepath.Dir(cwd)
		if parent == cwd {
			break // reached the root without finding go.mod
		}
		cwd = parent
	}

	return "", os.ErrNotExist
}
