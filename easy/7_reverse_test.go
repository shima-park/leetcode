package easy

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	var tests = []struct {
		X        int
		Expected int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{math.MaxInt64, 0},
		{math.MinInt64, 0},
		{-2147483648, 0},
	}

	for _, test := range tests {
		actual := reverse(test.X)
		assert.Equal(t, test.Expected, actual)
	}
}
