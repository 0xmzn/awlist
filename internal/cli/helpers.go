package cli

import (
	"fmt"
	"os"
)

func GetDataFilePath(providedPath string) (string, error) {
	if providedPath != "" {
		if _, err := os.Stat(providedPath); os.IsNotExist(err) {
			return "", fmt.Errorf("data file specified by --data flag not found: %s", providedPath)
		}
		return providedPath, nil
	}

	if _, err := os.Stat("awesome.yaml"); err == nil {
		return "awesome.yaml", nil
	} else if !os.IsNotExist(err) {
		return "", fmt.Errorf("error checking for awesome.yaml: %w", err)
	}

	return "", fmt.Errorf("no data file specified and awesome.yaml not found in current directory")
}
