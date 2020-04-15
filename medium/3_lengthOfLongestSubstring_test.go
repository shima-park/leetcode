package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	var tests = []struct {
		Str      string
		Expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{" ", 1},
		{"dvdf", 3},
	}

	for _, test := range tests {
		actual := lengthOfLongestSubstring(test.Str)
		assert.Equal(t, test.Expected, actual)
	}
}
