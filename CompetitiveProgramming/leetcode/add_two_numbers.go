/*
// @Author: Vicktor Desrony
// @filename: add_two_numbers.go
*/
package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ans := &ListNode{}
	curr := ans
	carry := 0

	for l1 != nil || l2 != nil {
		x := 0
		y := 0

		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}

		sum := carry + x + y
		carry = sum / 10
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
	}

	if carry > 0 {
		curr.Next = &ListNode{Val: carry}
	}

	return ans.Next
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func NewListNodeFromSlice(nums []int) *ListNode {
	l := NewListNode(nums[0])
	l.LoadFromSlice(nums[1:])
	return l
}

func (l *ListNode) Add(val int) {
	curr := l
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = &ListNode{Val: val}
}

func (l *ListNode) LoadFromSlice(nums []int) {
	curr := l
	for _, n := range nums {
		curr.Next = &ListNode{Val: n}
		curr = curr.Next
	}
}

func (l *ListNode) Equals(other *ListNode) bool {
	curr := l
	otherCurr := other

	for curr != nil && otherCurr != nil {
		if curr.Val != otherCurr.Val {
			return false
		}

		curr = curr.Next
		otherCurr = otherCurr.Next
	}

	return curr == nil && otherCurr == nil
}
