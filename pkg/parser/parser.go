// parser is the package containing tools
// to extract and standardize data from banks
package parser

// BankFilePath represents the path to
// a document provided by the bank
type BankFilePath string

// Interface that must be implemented
// by any data extractor, ie, the
// bank or file type specific parsers
type Interface interface {

	// GetEntries receives reads the
	// file in the BankFilePath and
	// parse its content to []Entry
	GetEntries(BankFilePath) ([]Entry, error)
}
