package ofx

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/thiagohdeplima/financial-control/pkg/parser"
)

func Test_GetEntries(t *testing.T) {
	cases := []struct {
		desc  string
		path  parser.BankFilePath
		expc  interface{}
		err   error
		setup func() *Parser
	}{
		{
			"returns error when file doesn't exists",
			"inexistent",
			[]parser.Entry{},
			&parser.FileNotFoundError{Path: "inexistent"},
			NewParser,
		},

		{
			"returns error when file content is invalid",
			"/tmp/invalid.content.file.ofx",
			[]parser.Entry{},
			&parser.InvalidBankFileError{
				Path: "/tmp/invalid.content.file.ofx",
				Err:  "Missing xml processing instruction",
			},
			func() *Parser {
				ioutil.WriteFile(
					"/tmp/invalid.content.file.ofx",
					[]byte("invalid content for an OFX"),
					0644,
				)

				return NewParser()
			},
		},

		{
			"returns Entries with no error when file content is valid",
			"/tmp/valid.file.ofx",
			[]parser.Entry{{
				ID:     "429167f0c125a9d91b9e7067f026f0b1",
				Date:   time.Date(2023, time.January, 2, 0, 0, 0, 0, time.FixedZone("GMT", -10800)),
				Desc:   "TESTING",
				Amount: -1000.1,
			}},
			nil,
			func() *Parser {
				content := []byte(`OFXHEADER:100
					DATA:OFXSGML
					VERSION:102
					SECURITY:NONE
					ENCODING:USASCII
					CHARSET:1252
					COMPRESSION:NONE
					OLDFILEUID:NONE
					NEWFILEUID:NONE
					<OFX>
						<SIGNONMSGSRSV1>
							<SONRS>
								<STATUS>
									<CODE>0
									<SEVERITY>INFO
								</STATUS>
								<DTSERVER>20221231000000[-3:GMT]
								<LANGUAGE>ENG
								<FI>
									<ORG>SANTANDER
									<FID>SANTANDER
								</FI>
							</SONRS>
						</SIGNONMSGSRSV1>
						<BANKMSGSRSV1>
							<STMTTRNRS>
								<TRNUID>1
								<STATUS>
									<CODE>0
									<SEVERITY>INFO
								</STATUS>
								<STMTRS>
									<CURDEF>BRL				
									<BANKACCTFROM>
										<BANKID>0001
										<ACCTID>00001
										<ACCTTYPE>CHECKING
									</BANKACCTFROM>
									<BANKTRANLIST>
										<DTSTART>20221224000000[-3:GMT]
										<DTEND>20221231000000[-3:GMT]
											<STMTTRN>
												<TRNTYPE>OTHER
												<DTPOSTED>20230102000000[-3:GMT]
												<TRNAMT>-1000.10
												<FITID>000000
												<CHECKNUM>000000
												<PAYEEID>0
												<MEMO>TESTING
											</STMTTRN>
										</BANKTRANLIST>
									<LEDGERBAL>
										<BALAMT>-2000.00
										<DTASOF>20221231000000[-3:GMT]
									</LEDGERBAL>
								</STMTRS>
							</STMTTRNRS>
						</BANKMSGSRSV1>
					</OFX>
				`)

				ioutil.WriteFile(
					"/tmp/valid.file.ofx",
					content,
					0644,
				)

				return NewParser()
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			prs := tc.setup()
			ret, err := prs.GetEntries(tc.path)

			assert.Equal(t, tc.expc, ret)
			assert.Equal(t, tc.err, err)
		})
	}
}
