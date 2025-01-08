package medium_test

import (
	"testing"

	"github.com/vldcration/go-ressources/CompetitiveProgramming/leetcode/problem/medium"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "Test1",
			nums:     []int{-1, 2, 1, -4},
			target:   1,
			expected: 2,
		},
		{
			name:     "Test1",
			nums:     []int{0, 0, 0},
			target:   1,
			expected: 0,
		},
	}

	for _, tt := range tests {
		if got := medium.ThreeSumClosestOptimized(tt.nums, tt.target); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}
