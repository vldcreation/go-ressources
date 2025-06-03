package easy_test

import (
	"testing"

	"github.com/vldcreation/go-ressources/CompetitiveProgramming/leetcode/problem/easy"
	"github.com/vldcreation/go-ressources/util"
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

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected bool
	}{
		{
			name:     "Test1",
			nums:     []int{1, 2, 3, 1},
			k:        3,
			expected: true,
		},
		{
			name:     "Test2",
			nums:     []int{1, 0, 1, 1},
			k:        1,
			expected: true,
		},
		{
			name:     "Test3",
			nums:     []int{1, 2, 3, 1, 2, 3},
			k:        2,
			expected: false,
		},
		{
			name:     "Test4",
			nums:     []int{1, 0, 0, 1, 3, 4, 1},
			k:        3,
			expected: true,
		},
	}

	for _, tt := range tests {
		if got := easy.ContainsNearbyDuplicate(tt.nums, tt.k); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestFindMaxAverage(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected float64
	}{
		{
			name:     "Test1",
			nums:     []int{1, 12, -5, -6, 50, 3},
			k:        4,
			expected: 12.75,
		},
		{
			name:     "Test2",
			nums:     []int{5},
			k:        1,
			expected: 5.0,
		},
	}

	for _, tt := range tests {
		if got := easy.FindMaxAverage(tt.nums, tt.k); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestFindLHS(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Test1",
			nums:     []int{1, 3, 2, 2, 5, 2, 3, 7},
			expected: 5,
		},
		{
			name:     "Test2",
			nums:     []int{1, 2, 3, 4},
			expected: 2,
		},
		{
			name:     "Test3",
			nums:     []int{1, 1, 1, 1},
			expected: 0,
		},
	}

	for _, tt := range tests {
		if got := easy.FindLHS(tt.nums); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestDecrypt(t *testing.T) {
	tests := []struct {
		name     string
		code     []int
		k        int
		expected []int
	}{
		// {
		// 	name:     "Test1",
		// 	code:     []int{5, 7, 1, 4},
		// 	k:        3,
		// 	expected: []int{12, 10, 16, 13},
		// },
		// {
		// 	name:     "Test2",
		// 	code:     []int{1, 2, 3, 4},
		// 	k:        0,
		// 	expected: []int{0, 0, 0, 0},
		// },
		{
			name:     "Test3",
			code:     []int{2, 4, 9, 3},
			k:        -2,
			expected: []int{12, 5, 6, 13},
		},
	}

	for _, tt := range tests {
		if got := easy.Decrypt(tt.code, tt.k); !util.CompareSliceInt(got, tt.expected) {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestLongestNiceSubstring(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name:     "Test1",
			s:        "YazaAay",
			expected: "aAa",
		},
		{
			name:     "Test2",
			s:        "Bb",
			expected: "Bb",
		},
		{
			name:     "Test3",
			s:        "c",
			expected: "",
		},
	}
	for _, tt := range tests {
		if got := easy.LongestNiceSubstring(tt.s); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestCountGoodSubstrings(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Test1",
			s:        "xyzzaz",
			expected: 1,
		},
		{
			name:     "Test2",
			s:        "aababcabc",
			expected: 4,
		},
	}

	for _, tt := range tests {
		if got := easy.CountGoodSubstrings(tt.s); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestMinimumDifference(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Test1",
			nums:     []int{90},
			k:        1,
			expected: 0,
		},
		{
			name:     "Test2",
			nums:     []int{9, 4, 1, 7},
			k:        2,
			expected: 2,
		},
	}

	for _, tt := range tests {
		if got := easy.MinimumDifference(tt.nums, tt.k); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestMinimumRecolors(t *testing.T) {
	tests := []struct {
		name     string
		blocks   string
		k        int
		expected int
	}{
		{
			name:     "Test1",
			blocks:   "WBBWWBB",
			k:        7,
			expected: 3,
		},
		{
			name:     "Test2",
			blocks:   "WBWBBBW",
			k:        2,
			expected: 0,
		},
	}

	for _, tt := range tests {
		if got := easy.MinimumRecolors(tt.blocks, tt.k); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}
