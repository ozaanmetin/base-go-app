package environment

import (
	"fmt"
	"path/filepath"
	"runtime"
)

func GetBaseDir() (string, error) {
	// Get the current file's directory, not the working directory
	_, currentFilePath, _, ok := runtime.Caller(0)
	if !ok {
		return "", fmt.Errorf("unable to get current file path")
	}

	// Get the directory of the current file
	baseDir := filepath.Dir(currentFilePath)
	baseDir, err := filepath.Abs(filepath.Join(baseDir, "..", "..", ".."))
	if err != nil {
		return "", err
	}
	return baseDir, nil
}
