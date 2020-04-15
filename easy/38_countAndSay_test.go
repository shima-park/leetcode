package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountAndSay(t *testing.T) {
	var tests = []struct {
		N        int
		Expected string
	}{
		{1, "1"},
		{4, "1211"},
	}

	for _, test := range tests {
		actual := countAndSay(test.N)
		assert.Equal(t, test.Expected, actual)
	}
}
