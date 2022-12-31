// parser is the package containing tools
// to extract and standardize data from banks
package parser

import (
	"time"
)

// BankFilePath represents the path to
// a document provided by the bank
type BankFilePath string

// Entry represents a financial
// entry extracted from the account
type Entry struct {
	ID     string
	Date   time.Time
	Desc   string
	Amount int
}

// Interface that must be implemented
// by any data extractor, ie, the
// bank specific parsers
type Interface interface {

	// GetEntries is where the entries
	// is extracted from the BankFilePath
	GetEntries(BankFilePath) []Entry
}
