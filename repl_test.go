package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "hello  world",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello  World",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		expectedLen := len(c.expected)
		actualLen := len(actual)

		if actualLen != expectedLen {
			t.Errorf("Length of actual: %v does match expected: %v", actualLen, expectedLen)
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("actual word: %v does not match expected word: %v", word, expectedWord)
			}
		}
	}
}
