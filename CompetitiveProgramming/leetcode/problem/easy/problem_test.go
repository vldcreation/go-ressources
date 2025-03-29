package easy_test

import (
	"testing"

	"github.com/vldcreation/go-ressources/CompetitiveProgramming/leetcode/problem/easy"
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
		name         string
		nums         []int
		expectedNums []int
		k            int
	}{
		{
			name:         "Test1",
			nums:         []int{1, 1, 2},
			expectedNums: []int{1, 2},
			k:            2,
		},
		{
			name:         "Test1",
			nums:         []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			expectedNums: []int{0, 1, 2, 3, 4},
			k:            5,
		},
	}

	for _, tt := range tests {
		if got := easy.RemoveDuplicatesFromSortedArray(tt.nums); got != tt.k {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expectedNums, got)
		}

		for i := 0; i < tt.k; i++ {
			if tt.nums[i] != tt.expectedNums[i] {
				t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expectedNums, tt.nums)
			}
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

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name     string
		head     *easy.ListNode
		expected bool
	}{
		{
			name:     "Test1",
			head:     easy.NewListNodeFromSlice([]int{3, 2, 0, -4}),
			expected: true,
		},
		{
			name:     "Test2",
			head:     easy.NewListNodeFromSlice([]int{1}),
			expected: false,
		},
	}

	for _, tt := range tests {
		if got := easy.HasCycle(tt.head); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name         string
		nums         []int
		expectedNums []int
		val          int
		k            int
	}{
		{
			name:         "Test1",
			nums:         []int{3, 2, 2, 3},
			expectedNums: []int{2, 2},
			val:          3,
			k:            2,
		},
		{
			name:         "Test2",
			nums:         []int{0, 1, 2, 2, 3, 0, 4, 2},
			expectedNums: []int{0, 1, 4, 0, 3},
			val:          2,
			k:            5,
		},
	}

	for _, tt := range tests {
		if got := easy.RemoveElement(tt.nums, tt.val); got != tt.k {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expectedNums, got)
		}

		for i := 0; i < tt.k; i++ {
			if tt.nums[i] != tt.expectedNums[i] {
				t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expectedNums, tt.nums)
			}
		}
	}
}
