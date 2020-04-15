package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	var tests = []struct {
		Str      string
		NumRows  int
		Expected string
	}{
		{
			"LEETCODEISHIRING",
			3,
			`LCIRETOESIIGEDHN`,
		},
		{
			"AB",
			1,
			`AB`,
		},
	}

	for _, test := range tests {
		actual := convert(test.Str, test.NumRows)
		assert.Equal(t, test.Expected, actual)
	}
}
