package leetcode_test

import (
	"testing"

	"github.com/vldcration/go-ressources/CompetitiveProgramming/leetcode"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1, l2, expected *leetcode.ListNode
	}{
		{
			l1:       leetcode.NewListNodeFromSlice([]int{9, 9, 9, 9, 9, 9, 9}),
			l2:       leetcode.NewListNodeFromSlice([]int{9, 9, 9, 9}),
			expected: leetcode.NewListNodeFromSlice([]int{8, 9, 9, 9, 0, 0, 0, 1}),
		},
		{
			l1:       leetcode.NewListNodeFromSlice([]int{2, 4, 3}),
			l2:       leetcode.NewListNodeFromSlice([]int{5, 6, 4}),
			expected: leetcode.NewListNodeFromSlice([]int{7, 0, 8}),
		},
	}

	for _, test := range tests {
		if got := leetcode.AddTwoNumbers(test.l1, test.l2); !test.expected.Equals(got) {
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
		if got := leetcode.LengthOfLongestSubstring(tt.test); got != tt.expected {
			t.Errorf("failed on test (%s): Expected %v, but got %v", tt.name, tt.expected, got)
		}
	}
}
