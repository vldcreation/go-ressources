package medium_test

import (
	"testing"

	"github.com/vldcreation/go-ressources/CompetitiveProgramming/leetcode/problem/medium"
	"github.com/vldcreation/go-ressources/util"
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

func TestFindMaxConsecutiveOnesI(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		k        int
		expected int
	}{
		{
			name:     "Test1",
			nums:     []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0},
			k:        2,
			expected: 6,
		},
		{
			name:     "Test1",
			nums:     []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1},
			k:        3,
			expected: 10,
		},
	}

	for _, tt := range tests {
		if got := medium.FindMaxConsecutiveOnesIII(tt.nums, tt.k); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestXorAllPairingNums(t *testing.T) {
	tests := []struct {
		name         string
		nums1, nums2 []int
		expected     int
	}{
		{
			name:     "Test1",
			nums1:    []int{2, 1, 3},
			nums2:    []int{10, 2, 5, 0},
			expected: 13,
		},
		{
			name:     "Test1",
			nums1:    []int{1, 2},
			nums2:    []int{3, 4},
			expected: 0,
		},
	}

	for _, tt := range tests {
		if got := medium.XorAllPairingNums(tt.nums1, tt.nums2); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestLetterCombinations(t *testing.T) {
	tests := []struct {
		name     string
		digits   string
		expected []string
	}{
		{
			name:     "Test1",
			digits:   "23",
			expected: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"},
		},
		{
			name:     "Test2",
			digits:   "",
			expected: []string{},
		},
		{
			name:     "Test3",
			digits:   "2",
			expected: []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		if got := medium.LetterCombinations(tt.digits); !util.CompareSliceString(tt.expected, got) {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestSumRange2D(t *testing.T) {

	expected := []*int{nil, medium.IntToPtr(8), medium.IntToPtr(11), medium.IntToPtr(12)}
	numArray := medium.Constructor([][]int{{3, 0, 1, 4, 2}, {5, 6, 3, 2, 1}, {1, 2, 0, 1, 5}, {4, 1, 0, 1, 7}, {1, 0, 3, 0, 5}})

	for i, v := range []struct {
		row1, col1, row2, col2 int
		expected               *int
	}{
		{2, 1, 4, 3, expected[1]},
		{1, 1, 2, 2, expected[2]},
		{1, 2, 2, 4, expected[3]},
	} {
		if got := numArray.SumRegion(v.row1, v.col1, v.row2, v.col2); got != *v.expected {
			t.Errorf("failed on test %d: Expected %v, but got %v", i, *v.expected, got)
		}
	}
}

func TestQueryResults(t *testing.T) {
	tests := []struct {
		name     string
		limit    int
		queries  [][]int
		expected []int
	}{
		{
			name:     "Test1",
			limit:    4,
			queries:  [][]int{{1, 4}, {2, 5}, {1, 3}, {3, 4}},
			expected: []int{1, 2, 2, 3},
		},
		{
			name:     "Test2",
			limit:    4,
			queries:  [][]int{{0, 1}, {1, 2}, {2, 2}, {3, 4}, {4, 5}},
			expected: []int{1, 2, 2, 3, 4},
		},
		{
			name:     "Test3",
			limit:    1,
			queries:  [][]int{{0, 1}, {0, 4}, {1, 2}, {1, 5}, {1, 4}},
			expected: []int{1, 1, 2, 2, 1},
		},
	}

	for _, tt := range tests {
		if got := medium.QueryResults(tt.limit, tt.queries); !util.CompareSliceInt(tt.expected, got) {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		name     string
		target   int
		nums     []int
		expected int
	}{
		{
			name:     "Test1",
			target:   7,
			nums:     []int{2, 3, 1, 2, 4, 3},
			expected: 2,
		},
		{
			name:     "Test2",
			target:   4,
			nums:     []int{1, 4, 4},
			expected: 1,
		},
		{
			name:     "Test3",
			target:   11,
			nums:     []int{1, 1, 1, 1, 1, 1, 1, 1},
			expected: 0,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := medium.MinSubArrayLen(test.target, test.nums)
			if result != test.expected {
				t.Errorf("%s - Expected %d, got %d", test.name, test.expected, result)
			}
		})
	}
}
