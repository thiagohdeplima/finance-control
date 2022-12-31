package ofx

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/aclindsa/ofxgo"
	"github.com/thiagohdeplima/financial-control/pkg/parser"
)

type Parser struct{}

func NewParser() *Parser {
	return &Parser{}
}

func (*Parser) GetEntries(path parser.BankFilePath) ([]parser.Entry, error) {
	if !fileExists(path) {
		return []parser.Entry{}, &parser.FileNotFoundError{Path: path}
	}

	resp, err := readBankFile(path)
	if err != nil {
		return []parser.Entry{}, &parser.InvalidBankFileError{Path: path, Err: err.Error()}
	}

	return getEntriesFromResponse(resp)
}

func fileExists(path parser.BankFilePath) bool {
	if _, err := os.Stat(string(path)); err == nil {
		return true
	} else {
		return false
	}
}

func readBankFile(path parser.BankFilePath) (*ofxgo.Response, error) {
	file, err := os.Open(string(path))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return ofxgo.ParseResponse(file)
}

func getEntriesFromResponse(resp *ofxgo.Response) ([]parser.Entry, error) {
	var entries = []parser.Entry{}

	if stmt, ok := resp.Bank[0].(*ofxgo.StatementResponse); ok {
		for _, tran := range stmt.BankTranList.Transactions {
			amount, _ := tran.TrnAmt.Float32()

			entries = append(entries, parser.Entry{
				ID:     getEntryID(tran),
				Amount: amount,
				Desc:   tran.Memo.String(),
				Date:   tran.DtPosted.Time,
			})
		}
	}

	return entries, nil
}

func getEntryID(tran ofxgo.Transaction) string {
	md5ID := md5.Sum([]byte(
		fmt.Sprintf("%s%s%s%s%s%s",
			tran.TrnType.String(),
			tran.Memo.String(),
			tran.DtPosted.String(),
			tran.TrnType.String(),
			tran.FiTID.String(),
			tran.CheckNum.String(),
		),
	))

	return hex.EncodeToString(md5ID[:])
}