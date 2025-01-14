package ifglife_test

import (
	"testing"

	"github.com/vldcration/go-ressources/CompetitiveProgramming/interview/ifglife"
)

func TestMinimumFlip(t *testing.T) {
	tests := []struct {
		name     string
		target   string
		expected int
	}{
		{
			name:     "test1",
			target:   "0011",
			expected: 1,
		},
		{
			name:     "test2",
			target:   "01011",
			expected: 3,
		},
	}

	for _, tt := range tests {
		if got := ifglife.MinimumFlips(tt.target); got != tt.expected {
			t.Errorf("expected: %+v but got: %+v\n", tt.expected, got)
		}
	}
}

func TestFindShortestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		target   string
		expected int
	}{
		{
			name:     "test1",
			target:   "xabbcacpqr",
			expected: 3,
		},
		{
			name:     "test2",
			target:   "abcb",
			expected: 1,
		},
		{
			name:     "test2",
			target:   "abc",
			expected: 0,
		},
	}

	for _, tt := range tests {
		if got := ifglife.FindShortestSubstring(tt.target); got != tt.expected {
			t.Errorf("expected: %+v but got: %+v\n", tt.expected, got)
		}
	}
}
