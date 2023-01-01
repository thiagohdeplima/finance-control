package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/thiagohdeplima/financial-control/pkg/parser"
)

type MockParser struct {
	mock.Mock
}

func (m *MockParser) GetEntries(YAMLFile parser.BankFilePath) ([]parser.Entry, error) {
	args := m.Called(YAMLFile)

	return args.Get(0).([]parser.Entry), args.Error(1)

}

func Test_ParseBankFile(t *testing.T) {
	cases := []struct {
		Desc       string
		ExpError   error
		ExpEntries []parser.Entry
	}{
		{
			"returns entries with no error when is parser returns entries with no error",
			nil,
			[]parser.Entry{
				{ID: "ID", Date: time.Now(), Desc: "DESCRIPTION", Amount: 1000.1},
			},
		},

		{
			"returns error when parser return error",
			errors.New("it's a testing"),
			[]parser.Entry{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Desc, func(t *testing.T) {
			mparser := new(MockParser)
			usecase := NewExtractEntriesFromBankFile(mparser)
			ymlPath := parser.BankFilePath("/dev/null")

			mparser.On("GetEntries", ymlPath).Return(tc.ExpEntries, tc.ExpError)

			actual, err := usecase.ExtractEntriesFromBankFile(ymlPath)

			assert.Equal(t, tc.ExpEntries, actual)
			assert.Equal(t, tc.ExpError, err)
		})
	}

	t.Run("returns the result of parser", func(t *testing.T) {

	})
}
