package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyAtoi(t *testing.T) {
	var tests = []struct {
		Str      string
		Expected int
	}{
		{"42", 42},
		{"   -42", -42},
		{"4193 with words", 4193},
		{"words and 987", 0},
		{"-9128347233", -2147483648},
		{"3.14159", 3},
		{"+1", 1},
		{"+-2", 0},
		{"   +0 123", 0},
		{`".1"`, 0},
		{`.1`, 0},
		{`-+1`, 0},
		{`0-1`, 0},
	}

	for _, test := range tests {
		actual := myAtoi(test.Str)
		fmt.Println(test.Str, test.Expected, actual)
		assert.Equal(t, test.Expected, actual, "Test Case: %s", test.Str)
	}
}
