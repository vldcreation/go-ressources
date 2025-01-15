package easy_test

import (
	"testing"

	"github.com/vldcration/go-ressources/CompetitiveProgramming/leetcode/problem/easy"
)

func TestFindMaxConsecutiveOnesI(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Test1",
			nums:     []int{1, 1, 0, 1, 1, 1},
			expected: 3,
		},
		{
			name:     "Test1",
			nums:     []int{1, 0, 1, 1, 0, 1},
			expected: 2,
		},
	}

	for _, tt := range tests {
		if got := easy.FindMaxConsecutiveOnes(tt.nums); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestRemoveDuplicatesFromSortedArray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Test1",
			nums:     []int{1, 1, 2},
			expected: 2,
		},
		{
			name:     "Test1",
			nums:     []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expected: 5,
		},
	}

	for _, tt := range tests {
		if got := easy.RemoveDuplicatesFromSortedArray(tt.nums); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}
