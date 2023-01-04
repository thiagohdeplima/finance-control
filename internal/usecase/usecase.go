package usecase

import (
	"github.com/thiagohdeplima/financial-control/pkg/parser"
)

// ExtractEntriesFromBankFile parses the content of
// parser.BankFilePath into []parser.Entry
type ExtractEntriesFromBankFile interface {
	Run(YAML parser.BankFilePath) ([]parser.Entry, error)
}
