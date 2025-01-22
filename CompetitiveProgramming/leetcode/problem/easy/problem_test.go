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

func intPtr(v int) *int {
	return &v
}
func TestSumRange(t *testing.T) {

	expected := []*int{nil, intPtr(1), intPtr(-1), intPtr(-3)}
	numArray := easy.Constructor([]int{-2, 0, 3, -5, 2, -1})

	for i, v := range []struct {
		left     int
		right    int
		expected *int
	}{
		{0, 2, expected[1]},
		{2, 5, expected[2]},
		{0, 5, expected[3]},
	} {
		if got := numArray.SumRange(v.left, v.right); got != *v.expected {
			t.Errorf("failed on test %d: Expected %v, but got %v", i, *v.expected, got)
		}
	}
}
