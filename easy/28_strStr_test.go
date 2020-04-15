package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrStr(t *testing.T) {
	var tests = []struct {
		Haystack string
		Needle   string
		Expected int
	}{
		{"mississippi", "issipi", -1},
	}

	for _, test := range tests {
		actual := strStr(test.Haystack, test.Needle)
		assert.Equal(t, test.Expected, actual)
	}
}
