package parser

import (
	"fmt"
)

type InvalidBankFileError struct {
	Path BankFilePath
	Err  string
}

func (e *InvalidBankFileError) Error() string {
	return fmt.Sprintf("the file %s is invalid: `%s`", e.Path, e.Err)
}
