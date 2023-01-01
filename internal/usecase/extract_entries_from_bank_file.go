package usecase

import (
	"github.com/thiagohdeplima/financial-control/pkg/parser"
)

type ExtractEntriesFromBankFileImpl struct {
	Parser parser.Interface
}

func NewExtractEntriesFromBankFile(parser parser.Interface) *ExtractEntriesFromBankFileImpl {
	return &ExtractEntriesFromBankFileImpl{Parser: parser}
}

func (p *ExtractEntriesFromBankFileImpl) ExtractEntriesFromBankFile(YAMLFile parser.BankFilePath) ([]parser.Entry, error) {
	return p.Parser.GetEntries(YAMLFile)
}
