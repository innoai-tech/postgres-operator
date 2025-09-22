package exec

import (
	"os"
	"path/filepath"
)

func WriteTempFile(filename string, data []byte) (string, error) {
	file := filepath.Join(os.TempDir(), filename)

	if err := os.WriteFile(file, data, 0o600); err != nil {
		return "", err
	}

	return file, nil
}
