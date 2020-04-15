package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	var tests = []struct {
		Str      string
		P        string
		Expected bool
	}{
		{"aa", "a", false},
		{"aa", "a*", true},
		{"ab", ".*", true},
		{"aab", "c*a*b", true},
		{"mississippi", "mis*is*p*.", false},
		{"ab", ".*c", false},
		{"aaa", "a*a", true},
		{"bbbba", ".*a*a", true},
		{"a", ".*..a*", false},
		{"ab", ".*..", true},
	}

	for _, test := range tests {
		actual := isMatch(test.Str, test.P)
		assert.Equal(t, test.Expected, actual,
			"Test Case:%s, Test Pattern %s", test.Str, test.P,
		)
	}
}
