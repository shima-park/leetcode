package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompressString(t *testing.T) {
	var tests = []struct {
		Str      string
		Expected string
	}{
		{"aabcccccaaa", "a2b1c5a3"},
		{"abbccd", "abbccd"},
	}

	for _, test := range tests {
		actual := compressString(test.Str)
		assert.Equal(t, test.Expected, actual,
			"Test Case:%s", test.Str,
		)
	}
}
