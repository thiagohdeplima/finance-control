package fileutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_FileExists(t *testing.T) {
	cases := []struct {
		Desc     string
		FilePath string
		Result   bool
	}{
		{"return true when file exists", "/dev/null", true},
		{"return false when file doesn't exists", "/tmp/file.doesnt.exists", false},
	}

	for _, tc := range cases {
		t.Run(tc.Desc, func(t *testing.T) {
			assert.Equal(t, tc.Result, FileExists(tc.FilePath))
		})
	}
}
