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
			input:    "   hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hElLo My NaMe Is",
			expected: []string{"hello", "my", "name", "is"},
		},
		{
			input:    "STAR TREK IS AWESOME",
			expected: []string{"star", "trek", "is", "awesome"},
		},
		// add more cases
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected) {
			t.Errorf("Number of elements returned: %v did not pass expected: %v", actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("Word: %v does not match expected word: %v", word, expectedWord)
			}
		}
	}
}
