package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " hello  world ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hi Hello Hey",
			expected: []string{"hi", "hello", "hey"},
		},
		// add more cases here

	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Fail. Expected: %s, Actual: %s", expectedWord, word)
			}
		}
	}
}
