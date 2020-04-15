package leetcode

import (
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	var tests = []struct {
		Str      string
		Expected []string
	}{
		{"bb", []string{"bb"}},
		{"aba", []string{"aba"}},
		{"abba", []string{"abba"}},
		{"babad", []string{"bab", "aba"}},
		{"cbbd", []string{"bb"}},
	}

	for _, test := range tests {
		actual := longestPalindrome(test.Str)
		var pass bool
		for _, expected := range test.Expected {
			if expected == actual {
				pass = true
			}
		}
		if !pass {
			t.Fatal("expected match any answer", test.Expected, "but got", actual)
		}
	}
}
