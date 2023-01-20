package utilities

import (
	"testing"
)

type FormatterTest struct {
	test_name string
	cue       []string
	expected  string
}

func TestFormmater(t *testing.T) {
	tests := []FormatterTest{
		{
			test_name: "More than one word",
			cue:       []string{"Hello", "There"},
			expected:  "Hello+There",
		},
		{
			test_name: "Just_one_word",
			cue:       []string{"Hello"},
			expected:  "Hello",
		},
	}

	for _, test := range tests {
		t.Run(test.test_name, func(t *testing.T) {
			if got := StringFormatter(test.cue); got != test.expected {
				t.Errorf("Wanted: %s, Got: %s", test.expected, got)
			}
		})
	}
}
