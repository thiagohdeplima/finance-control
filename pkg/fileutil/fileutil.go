package fileutil

import (
	"fmt"
	"os"
)

type FileNotFoundError struct {
	Path string
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("the file %s doesn't exists", e.Path)
}

func FileExists(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}

	return false
}

func FileExistsWithError(filePath string) error {
	if FileExists(filePath) {
		return nil
	}

	return &FileNotFoundError{filePath}
}
