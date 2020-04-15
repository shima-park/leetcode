package easy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	var tests = []struct {
		S        string
		Expected bool
	}{
		//{"()", true},
		//{"()[]{}", true},
		//{"(]", false},
		//{"([)]", false},
		//{"{[]}", true},
		//{"", true},
		{"(([]){})", true},
	}

	for _, test := range tests {
		actual := isValid(test.S)
		assert.Equal(t, test.Expected, actual,
			"Test case: %s", test.S,
		)
	}
}
