package medium_test

import (
	"testing"

	"github.com/vldcreation/go-ressources/CompetitiveProgramming/coderbyte/problem/medium"
)

func TestMinWindowSubstring(t *testing.T) {
	tests := []struct {
		name     string
		s        [2]string
		expected string
	}{
		{
			name:     "test 1",
			s:        [2]string{"ahffaksfajeeubsne", "jefaa"},
			expected: "aksfaje",
		},
		{
			name:     "test 2",
			s:        [2]string{"aaffhkksemckelloe", "fhea"},
			expected: "affhkkse",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := medium.MinWindowSubstring(test.s)
			if result != test.expected {
				t.Errorf("Expected %s, got %s", test.expected, result)
			}
		})
	}
}
