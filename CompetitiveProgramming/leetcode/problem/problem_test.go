package problem_test

import (
	"testing"

	"github.com/vldcreation/go-ressources/CompetitiveProgramming/interview"
	"github.com/vldcreation/go-ressources/CompetitiveProgramming/leetcode/problem/easy"
	"github.com/vldcreation/go-ressources/CompetitiveProgramming/leetcode/problem/medium"
	"github.com/vldcreation/go-ressources/util"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1, l2, expected *medium.ListNode
	}{
		{
			l1:       medium.NewListNodeFromSlice([]int{2, 4, 3}),
			l2:       medium.NewListNodeFromSlice([]int{5, 6, 4}),
			expected: medium.NewListNodeFromSlice([]int{7, 0, 8}),
		},
		{
			l1:       medium.NewListNodeFromSlice([]int{9, 9, 9, 9, 9, 9, 9}),
			l2:       medium.NewListNodeFromSlice([]int{9, 9, 9, 9}),
			expected: medium.NewListNodeFromSlice([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
	}

	for _, test := range tests {
		if got := medium.AddTwoNumbers(test.l1, test.l2); !test.expected.Equals(got) {
			t.Errorf("Expected %v, but got %v", test.expected, got)
		}
	}
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		name     string
		test     string
		expected int
	}{
		{
			name:     "test 1",
			test:     "abcabcbb",
			expected: 3,
		},
		{
			name:     "test 2",
			test:     "pwwkew",
			expected: 3,
		},
		{
			name:     "test 3",
			test:     "bbbbb",
			expected: 1,
		},
		{
			name:     "test 4",
			test:     "au",
			expected: 2,
		},
		{
			name:     "test 5",
			test:     "dvdf",
			expected: 3,
		},
		{
			name:     "test 6",
			test:     "nfpdmpi",
			expected: 5,
		},
	}

	for _, tt := range tests {
		if got := medium.LengthOfLongestSubstring(tt.test); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestReverseInt(t *testing.T) {
	tests := []struct {
		name     string
		x        int
		expected int
	}{
		{
			name:     "Test 1",
			x:        121,
			expected: 121,
		},
		{
			name:     "Test 2",
			x:        -123,
			expected: -321,
		},
		{
			name:     "Test 3",
			x:        120,
			expected: 21,
		},
		{
			name:     "Test 4",
			x:        1234123412341234123,
			expected: 0,
		},
		{
			name:     "Test 5",
			x:        1534236469,
			expected: 0,
		},
	}

	for _, tt := range tests {
		if got := medium.ReverseInt(tt.x); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestZigzagConversion(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		numRows  int
		expected string
	}{
		{
			name:     "Test 1",
			s:        "PAYPALISHIRING",
			numRows:  2,
			expected: "PYAIHRNAPLSIIG",
		},
		{
			name:     "Test 2",
			s:        "PA",
			numRows:  2,
			expected: "PA",
		},
		{
			name:     "Test 3",
			s:        "PAYPALISHIRING",
			numRows:  3,
			expected: "PAHNAPLSIIGYIR",
		},
		{
			name:     "Test 4",
			s:        "PAYPALISHIRING",
			numRows:  4,
			expected: "PINALSIGYAHRPI",
		},
		{
			name:     "Test 5",
			s:        "ABC",
			numRows:  2,
			expected: "ACB",
		},
		{
			name:     "Test 5",
			s:        "ABC",
			numRows:  3,
			expected: "ABC",
		},
		{
			name:     "Test 5",
			s:        "ABCDE",
			numRows:  4,
			expected: "ABCED",
		},
	}

	for _, tt := range tests {
		if got := medium.ZigzagConversion__BEST(tt.s, tt.numRows); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
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
		result := medium.LongestPalindrome(tc.str)

		if !util.CheckStringInSlice(result, tc.expected) {
			t.Errorf("%s failed, Expected one of %v, got %s", tc.name, tc.expected, result)
		}
	}
}

func TestIntersect(t *testing.T) {
	tests := []struct {
		name     string
		nums1    []int
		nums2    []int
		expected [][]int
	}{
		{
			name:  "TC1",
			nums1: []int{1, 2, 2, 1},
			nums2: []int{2, 2},
			expected: [][]int{
				{2, 2},
			},
		},
		{
			name:  "TC2",
			nums1: []int{4, 9, 5},
			nums2: []int{9, 4, 9, 8, 4},
			expected: [][]int{
				{4, 9},
			},
		},
		{
			name:  "TC3",
			nums1: []int{0, 2, 7, 7, 9},
			nums2: []int{0, 8, 4},
			expected: [][]int{
				{0},
			},
		},
		{
			name:  "TC4",
			nums1: []int{2, 1},
			nums2: []int{1, 1},
			expected: [][]int{
				{1},
			},
		},
		{
			name:  "TC5",
			nums1: []int{2, 1},
			nums2: []int{1, 2},
			expected: [][]int{
				{1, 2},
			},
		},
	}

	for _, tc := range tests {
		result := medium.Intersect(tc.nums1, tc.nums2)

		ok := false
		for _, v := range tc.expected {
			if util.CompareSliceInt(result, v) {
				ok = true
				break
			}
		}

		if !ok {
			t.Errorf("%s failed, Expected one of %v, got %v", tc.name, tc.expected, result)
		}
	}
}

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		expected int
	}{
		{
			name:     "Test1",
			s:        "1337c0d3",
			expected: 1337,
		},
		{
			name:     "Test2",
			s:        "   -042",
			expected: -42,
		},
		{
			name:     "Test3",
			s:        "0-1",
			expected: 0,
		},
		{
			name:     "Test4",
			s:        "-2147483647",
			expected: -2147483647,
		},
		{
			name:     "Test5",
			s:        "20000000000000000000",
			expected: 2147483647,
		},
		{
			name:     "Test6",
			s:        "words and 987",
			expected: 0,
		},
	}

	for _, tt := range tests {
		if got := medium.MyAtoi__GPT(tt.s); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Test1",
			input:    []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "Test2",
			input:    []int{1, 1},
			expected: 1,
		},
		{
			name:     "Test3",
			input:    []int{1, 1},
			expected: 1,
		},
	}

	for _, tt := range tests {
		if got := medium.MaxArea(tt.input); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestIntToRoman(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected string
	}{
		{
			name:     "Test1",
			input:    3749,
			expected: "MMMDCCXLIX",
		},
		{
			name:     "Test2",
			input:    58,
			expected: "LVIII",
		},
		{
			name:     "Test3",
			input:    1994,
			expected: "MCMXCIV",
		},
	}

	for _, tt := range tests {
		if got := medium.IntToRoman__GPT(tt.input); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected [][]int
	}{
		{
			name:     "Test1",
			input:    []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
	}

	for _, tt := range tests {
		if got := medium.ThreeSum(tt.input); !util.Compare2DSliceInt(got, tt.expected) {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestArrayChallanger(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected string
	}{
		{
			name:     "Test1",
			input:    []int{1, 42, 42, 180},
			expected: "42 _27p_2ds_a",
		},
		{
			name:     "Test2",
			input:    []int{4, 90},
			expected: "90 _7p3_dsq_",
		},
		// {
		// 	name:     "Test3",
		// 	input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		// 	expected: "2 _27p_2ds_a",
		// },
	}

	for _, tt := range tests {
		if got := interview.ArrayChallenge(tt.input); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestSearchChallenge(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test1",
			input:    "Hello apple pie",
			expected: "Hel_o7p_2ds_a",
		},
		{
			name:     "Test2",
			input:    "No words",
			expected: "-17_32d_qa",
		},
		// {
		// 	name:     "Test3",
		// 	input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		// 	expected: "2 _27p_2ds_a",
		// },
	}

	for _, tt := range tests {
		if got := interview.SearchChallenge(tt.input); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestMathChallenge(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Test1",
			input:    "6*(4/2)+3*1",
			expected: "15",
		},
		{
			name:     "Test2",
			input:    "100*2**4",
			expected: "1600",
		},
		// {
		// 	name:     "Test3",
		// 	input:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		// 	expected: "2 _27p_2ds_a",
		// },
	}

	for _, tt := range tests {
		if got := interview.MathChallenge(tt.input); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestIsPalindromeLinkedList(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{
			name:     "Test1",
			input:    []int{2, 4, 6, 4, 2},
			expected: true,
		},
		{
			name:     "Test2",
			input:    []int{1, 2, 2, 1},
			expected: true,
		},
		{
			name:     "Test3",
			input:    []int{1, 2},
			expected: false,
		},
	}

	for _, tt := range tests {
		if got := easy.IsPalindromeLinkedList(easy.MakeListFromSlice(tt.input)); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		name           string
		input1, input2 []int
		expected       []int
	}{
		{
			name:     "Test1",
			input1:   []int{1, 2, 4},
			input2:   []int{1, 3, 4},
			expected: []int{1, 1, 2, 3, 4, 4},
		},
		{
			name:     "Test2",
			input1:   []int{},
			input2:   []int{},
			expected: []int{},
		},
		{
			name:     "Test3",
			input1:   []int{},
			input2:   []int{1, 2},
			expected: []int{1, 2},
		},
	}

	for _, tt := range tests {
		if got := easy.MergeTwoLists(easy.MakeListFromSlice(tt.input1), easy.MakeListFromSlice(tt.input2)); !util.CompareSliceInt(easy.ExtractListToSlice(got), tt.expected) {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}
