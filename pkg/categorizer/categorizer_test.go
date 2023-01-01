package categorizer

import (
	"errors"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thiagohdeplima/financial-control/pkg/fileutil"
)

func Test_ParseRules(t *testing.T) {
	cases := []struct {
		Desc  string
		File  string
		Err   error
		Rules Rules
		Setup func()
	}{
		{
			"returns error when file doesn't exists",
			"inexistent",
			&fileutil.FileNotFoundError{Path: "inexistent"},
			Rules{},
			func() {},
		},
		{
			"returns error when file is invalid",
			"/tmp/valid.rules.yaml",
			errors.New("yaml: line 2: found character that cannot start any token"),
			Rules{},
			func() {
				content := []byte(`---
				rules:
				- pattern: STARBUCKS
					category: BREAKFAST
				- pattern: MC(.*)DONALDS
					category: LUNCH
				`)

				ioutil.WriteFile(
					"/tmp/valid.rules.yaml",
					content,
					0644,
				)
			},
		},
		{
			"returns parsed rules when is everything ok",
			"/tmp/valid.rules.yaml",
			nil,
			Rules{
				Rules: []Rule{
					{Pattern: "STARBUCKS", Category: "BREAKFAST"},
					{Pattern: "MC(.*)DONALDS", Category: "LUNCH"},
				},
			},
			func() {
				content := []byte(`---
rules:
- pattern: STARBUCKS
  category: BREAKFAST
- pattern: MC(.*)DONALDS
  category: LUNCH`)

				ioutil.WriteFile(
					"/tmp/valid.rules.yaml",
					content,
					0644,
				)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.Desc, func(t *testing.T) {
			tc.Setup()

			actual, err := ParseRules(tc.File)

			assert.Equal(t, tc.Err, err)
			assert.Equal(t, tc.Rules, actual)
		})
	}
}
