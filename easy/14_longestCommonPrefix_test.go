package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonPrefix(t *testing.T) {
	var tests = []struct {
		Strs     []string
		Expected string
	}{
		{[]string{"a"}, "a"},
		{[]string{"aa", "a"}, "a"},
	}

	for _, test := range tests {
		actual := longestCommonPrefix(test.Strs)
		assert.Equal(t, test.Expected, actual)
	}
}
