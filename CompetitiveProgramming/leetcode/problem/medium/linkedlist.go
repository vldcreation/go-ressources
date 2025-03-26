package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func MakeListFromSlice(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := new(ListNode)
	curr := l
	for _, n := range nums {
		curr.Next = &ListNode{Val: n}
		curr = curr.Next
	}

	return l.Next
}

func ExtractListToSlice(head *ListNode) []int {
	var nums []int
	for head != nil {
		nums = append(nums, head.Val)
		head = head.Next
	}
	return nums
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}

func NewListNodeFromSlice(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := NewListNode(nums[0])
	if len(nums) == 1 {
		return l
	}

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
