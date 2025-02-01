package oy_test

import (
	"testing"

	"github.com/vldcreation/go-ressources/CompetitiveProgramming/interview/oy"
)

func TestCase1(t *testing.T) {
	tests := []struct {
		name     string
		a, b, k  int
		expected int
	}{
		{
			name:     "test1",
			a:        1,
			b:        10,
			k:        3,
			expected: 3,
		},
		{
			name:     "test2",
			a:        8,
			b:        20,
			k:        4,
			expected: 4,
		},
	}

	for _, tt := range tests {
		if got := oy.Case1(tt.a, tt.b, tt.k); got != tt.expected {
			t.Errorf("%s expected: %+v but got: %+v\n", tt.name, tt.expected, got)
		}
	}
}
