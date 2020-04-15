package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	var tests = []struct {
		X        int
		Expected bool
	}{
		{121, true},
		{-121, false},
		{10, false},
		{111121111, true},
		{22222222, true},
		{1, true},
	}

	for _, test := range tests {
		actual := isPalindrome(test.X)
		assert.Equal(t, test.Expected, actual,
			"Test Case: %d", test.X,
		)
	}
}
