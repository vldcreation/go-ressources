package medium_test

import (
	"testing"

	"github.com/vldcration/go-ressources/CompetitiveProgramming/leetcode/problem/medium"
	"github.com/vldcration/go-ressources/util"
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

func Test4Sum(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected [][]int
	}{
		{
			name:     "Test1",
			nums:     []int{1, 0, -1, 0, -2, 2},
			target:   0,
			expected: [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}},
		},
		{
			name:     "Test1",
			nums:     []int{2, 2, 2, 2, 2},
			target:   8,
			expected: [][]int{{2, 2, 2, 2}},
		},
	}

	for _, tt := range tests {
		if got := medium.FourSum(tt.nums, tt.target); !util.Compare2DSliceInt(tt.expected, got) {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}
