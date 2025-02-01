package twopointer_test

import (
	"testing"

	"github.com/vldcreation/go-ressources/CompetitiveProgramming/pattern/twopointer"
)

func TestCheckIfOneSwapCanMakeEqual(t *testing.T) {
	tests := []struct {
		name     string
		s1, s2   string
		expected bool
	}{
		{
			name:     "test 1",
			s1:       "ab",
			s2:       "ba",
			expected: true,
		},
		{
			name:     "test 2",
			s1:       "bank",
			s2:       "kanb",
			expected: true,
		},
	}

	for _, tt := range tests {
		if got := twopointer.CheckIfOneSwapCanMakeEqual(tt.s1, tt.s2); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}
