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
			AssertEqual(t, StringFormatter(test.cue), test.expected)
		})
	}
}

func AssertEqual[T comparable](t *testing.T, got, wanted T) {
	t.Helper()

	if got != wanted {
		t.Errorf("Wanted: %v, Got: %v", wanted, got)
	}
}
