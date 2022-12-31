package parser

import (
	"fmt"
)

type FileNotFoundError struct {
	Path BankFilePath
}

type InvalidBankFileError struct {
	Path BankFilePath
	Err  string
}

func (e *FileNotFoundError) Error() string {
	return fmt.Sprintf("the file %s doesn't exists", e.Path)
}

func (e *InvalidBankFileError) Error() string {
	return fmt.Sprintf("the file %s is invalid: `%s`", e.Path, e.Err)
}
