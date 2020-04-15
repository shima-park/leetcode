package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLastRemaining(t *testing.T) {
	var tests = []struct {
		N        int
		M        int
		Expected int
	}{
		{5, 3, 3},
		{10, 17, 2},
	}

	for _, test := range tests {
		actual := lastRemaining(test.N, test.M)
		assert.Equal(t, test.Expected, actual)
	}
}
