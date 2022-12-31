package parser

import (
	"fmt"
)

type FileNotFoundError struct {
	Path BankFilePath
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("the file %s doesn't exists", e.Path)
}
