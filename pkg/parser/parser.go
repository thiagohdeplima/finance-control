// parser is the package containing tools
// to extract and standardize data from banks
package parser

// BankFilePath represents the path to
// a document provided by the bank
type BankFilePath string

// Interface that must be implemented
// by any data extractor, ie, the
// bank specific parsers
type Interface interface {

	// GetEntries is where the entries
	// is extracted from the BankFilePath
	GetEntries(BankFilePath) ([]Entry, error)
}
