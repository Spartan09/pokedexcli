package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct{
		input string
		expected []string
	} {
		{
			input: "hello world!",
			expected: []string{"hello", "world!"},
		},
		{
			input: "HELLO world!",
			expected: []string{"hello", "world!"},
		},
	}

	for _, cs := range cases {
		acutal := cleanInput(cs.input)
		expected := cs.expected

		if len(acutal) != len(expected) {
			t.Errorf("length not equal: %v 'VS' %v", len(acutal), len(expected))
			continue
		}

		for i := range acutal {
			acutalWord := acutal[i]
			expectedWord := cs.expected[i]
			if acutalWord != expectedWord {
				t.Errorf("'%v' not equal to '%v'", acutalWord, expectedWord)
			}
		}

	}

}
