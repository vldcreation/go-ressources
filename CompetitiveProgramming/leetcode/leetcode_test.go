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
