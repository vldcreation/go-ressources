package easy

func HasCycle(head *ListNode) bool {
	slow := head
	if slow == nil {
		return false
	}

	fast := slow.Next

	for fast != nil && fast.Next != nil {
		if fast == slow {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}
