package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "   ",
			expected: []string{},
		},
		{
			input:    "Hello, World!",
			expected: []string{"hello,", "world!"},
		},
		{
			input:    "  Hello   World  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, tc := range cases {
		actual := cleanInput(tc.input)
		if len(actual) != len(tc.expected) {
			t.Errorf("Lengths don't match: Expected %v, got %v", tc.expected, actual)
			continue
		}
		for i, v := range tc.expected {
			if actual[i] != v {
				t.Errorf("Expected %v at index %d, got %v", v, i, actual[i])
			}
		}
	}
}
