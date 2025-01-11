package easy_test

import (
	"testing"

	"github.com/vldcration/go-ressources/CompetitiveProgramming/leetcode/problem/easy"
)

func Test4Sum(t *testing.T) {
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
