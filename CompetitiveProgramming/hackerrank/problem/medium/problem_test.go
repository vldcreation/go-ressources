package medium_test

import (
	"testing"

	"github.com/vldcration/go-ressources/CompetitiveProgramming/hackerrank/problem/medium"
)

func TestFormingMagicSquare(t *testing.T) {
	tests := []struct {
		name     string
		nums     [][]int32
		expected int32
	}{
		{
			name:     "Test1",
			nums:     [][]int32{{4, 9, 2}, {3, 5, 7}, {8, 1, 5}},
			expected: 1,
		},
		{
			name:     "Test1",
			nums:     [][]int32{{4, 8, 2}, {4, 5, 7}, {6, 1, 6}},
			expected: 4,
		},
	}

	for _, tt := range tests {
		if got := medium.FormingMagicSquare(tt.nums); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}
