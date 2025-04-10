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

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "Test1",
			nums:     []int{1, 2, 3, 4},
			expected: []int{24, 12, 8, 6},
		},
		{
			name:     "Test2",
			nums:     []int{-1, 1, 0, -3, 3},
			expected: []int{0, 0, 9, 0, 0},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := medium.ProductExceptSelfBest(test.nums)
			if !util.CompareSliceInt(test.expected, result) {
				t.Errorf("%s - Expected %v, got %v", test.name, test.expected, result)
			}
		})
	}
}

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		expected []string
	}{
		{
			name:     "Test 1",
			str:      "babad",
			expected: []string{"bab", "aba"},
		},
		{
			name:     "Test 2",
			str:      "cbbd",
			expected: []string{"bb"},
		},
		{
			name:     "Test 3",
			str:      "ccc",
			expected: []string{"ccc"},
		},
	}

	for _, tc := range tests {
		result := medium.LongestPalindromeBest(tc.str)

		if !util.CheckStringInSlice(result, tc.expected) {
			t.Errorf("%s failed, Expected one of %v, got %s", tc.name, tc.expected, result)
		}
	}
}

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		name     string
		head     *medium.ListNode
		n        int
		expected *medium.ListNode
	}{
		{
			name:     "Test 1",
			head:     medium.NewListNodeFromSlice([]int{1, 2, 3, 4, 5}),
			n:        2,
			expected: medium.NewListNodeFromSlice([]int{1, 2, 3, 5}),
		},
		{
			name:     "Test 2",
			head:     medium.NewListNodeFromSlice([]int{1}),
			n:        1,
			expected: medium.NewListNodeFromSlice(nil),
		},
		{
			name:     "Test 3",
			head:     medium.NewListNodeFromSlice([]int{1, 2}),
			n:        1,
			expected: medium.NewListNodeFromSlice([]int{1}),
		},
	}

	for _, tc := range tests {
		result := medium.RemoveNthFromEnd(tc.head, tc.n)

		if !result.Equals(tc.expected) {
			t.Errorf("%s failed, Expected %v, got %v", tc.name, tc.expected, result)
		}
	}
}

func TestEqualSubstringWithinBudget(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		t        string
		maxCost  int
		expected int
	}{
		{
			name:     "Test 1",
			s:        "abcd",
			t:        "bcdf",
			maxCost:  3,
			expected: 3,
		},
		{
			name:     "Test 2",
			s:        "abcd",
			t:        "cdef",
			maxCost:  3,
			expected: 1,
		},
		{
			name:     "Test 3",
			s:        "abcd",
			t:        "acde",
			maxCost:  0,
			expected: 1,
		},
	}

	for _, tc := range tests {
		result := medium.EqualSubstring(tc.s, tc.t, tc.maxCost)

		if result != tc.expected {
			t.Errorf("%s failed, Expected %v, got %v", tc.name, tc.expected, result)
		}
	}
}

func TestBalancedStringWithinBudget(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Test 1",
			s:        "QWER",
			expected: 0,
		},
		{
			name:     "Test 2",
			s:        "QQWE",
			expected: 1,
		},
		{
			name:     "Test 3",
			s:        "QQQW",
			expected: 2,
		},
		{
			name:     "Test 4",
			s:        "WQWRQQQW",
			expected: 3,
		},
	}

	for _, tc := range tests {
		result := medium.BalancedString(tc.s)

		if result != tc.expected {
			t.Errorf("%s failed, Expected %v, got %v", tc.name, tc.expected, result)
		}
	}
}

func TestLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{
			name:     "Test1",
			s:        "aaabb",
			k:        3,
			expected: 3,
		},
		{
			name:     "Test2",
			s:        "ababbc",
			k:        2,
			expected: 5,
		},
	}

	for _, tt := range tests {
		if got := medium.LongestSubstring(tt.s, tt.k); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestLongestSubarray(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Test1",
			nums:     []int{1, 1, 0, 1},
			expected: 3,
		},
		{
			name:     "Test2",
			nums:     []int{0, 1, 1, 1, 0, 1, 1, 0, 1},
			expected: 5,
		},
		{
			name:     "Test3",
			nums:     []int{1, 1, 1},
			expected: 2,
		},
		{
			name:     "Test4",
			nums:     []int{1, 1, 0, 0, 1, 1, 1, 0, 1},
			expected: 4,
		},
	}

	for _, tt := range tests {
		if got := medium.LongestSubarray(tt.nums); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestFindAnagrams(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		p        string
		expected []int
	}{
		{
			name:     "Test1",
			s:        "cbaebabacd",
			p:        "abc",
			expected: []int{0, 6},
		},
		{
			name:     "Test2",
			s:        "abab",
			p:        "ab",
			expected: []int{0, 1, 2},
		},
	}

	for _, tt := range tests {
		if got := medium.FindAnagrams(tt.s, tt.p); !util.CompareSliceInt(tt.expected, got) {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestCharacterReplacement(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		k        int
		expected int
	}{
		{
			name:     "Test1",
			s:        "ABAB",
			k:        2,
			expected: 4,
		},
		{
			name:     "Test2",
			s:        "AABABBA",
			k:        1,
			expected: 4,
		},
		{
			name:     "Test3",
			s:        "ABBB",
			k:        2,
			expected: 4,
		},
	}

	for _, tt := range tests {
		if got := medium.CharacterReplacement(tt.s, tt.k); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestNumberOfArithmeticSlices(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{
			name:     "Test1",
			nums:     []int{1, 2, 3, 4},
			expected: 3,
		},
		{
			name:     "Test2",
			nums:     []int{1},
			expected: 0,
		},
	}

	for _, tt := range tests {
		if got := medium.NumberOfArithmeticSlices(tt.nums); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestFindClosestElements(t *testing.T) {
	tests := []struct {
		name     string
		arr      []int
		k        int
		x        int
		expected []int
	}{
		{
			name:     "Test1",
			arr:      []int{1, 2, 3, 4, 5},
			k:        4,
			x:        3,
			expected: []int{1, 2, 3, 4},
		},
		{
			name:     "Test2",
			arr:      []int{1, 1, 2, 3, 4, 5},
			k:        4,
			x:        -1,
			expected: []int{1, 1, 2, 3},
		},
	}

	for _, tt := range tests {
		if got := medium.FindClosestElements(tt.arr, tt.k, tt.x); !util.CompareSliceInt(tt.expected, got) {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}
