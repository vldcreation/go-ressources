package leetcode_test

import (
	"testing"

	"github.com/vldcration/go-ressources/CompetitiveProgramming/leetcode"
)

func TestMustLoadProblems(t *testing.T) {
	tests := []struct {
		name     string
		expected int
	}{
		{
			name:     "test 1",
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Logf("problems: %v", leetcode.NewLeetcode(leetcode.MustLoadProblems()))
		if got := len(leetcode.NewLeetcode(leetcode.MustLoadProblems()).Problems); got != test.expected {
			t.Errorf("Expected %v, but got %v", test.expected, got)
		}
	}
}
