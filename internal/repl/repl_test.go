package repl

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
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
	}

	for _, testcase := range cases {
		actual := CleanInput(testcase.input)

		expectedLen := len(testcase.expected)
		actualLen := len(actual)
		if expectedLen != actualLen {
			t.Errorf("length mismatch. Expected: %d | Actual: %d", expectedLen, actualLen)
			return
		}
		for i := range actual {
			word := actual[i]
			expectedWord := testcase.expected[i]
			if expectedWord != word {
				t.Errorf("word mismatch. Expected: %s | Actual: %s\n", expectedWord, word)
				return
			}
		}
	}
}
