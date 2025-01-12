package unittests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPositive(t *testing.T) {
	testCases := []struct {
		name     string
		intVal   int
		expected string
	}{
		{
			name:     "int is 0",
			intVal:   0,
			expected: "i is 0",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name,
			func(t *testing.T) {
				assert.Equal(t, testCase.expected, isPositive(testCase.intVal))
			})
	}

}
