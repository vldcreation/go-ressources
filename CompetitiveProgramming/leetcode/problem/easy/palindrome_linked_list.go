package easy

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverse(head *ListNode) *ListNode {
	var prev, curr *ListNode
	curr = head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func IsPalindromeLinkedList(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// Find the middle of the linked list
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// Reverse the second half of the linked list
	secondHalf := reverse(slow)

	// Copy the head of the reversed second half
	copyHeadSecondHalf := secondHalf

	// Compare the first and second halves
	for head != nil && secondHalf != nil {
		if head.Val != secondHalf.Val {
			return false
		}
		head = head.Next
		secondHalf = secondHalf.Next
	}

	// Revert the reversed second half back to original state
	reverse(copyHeadSecondHalf)

	// If both halves match, the linked list is a palindrome
	return head == nil || secondHalf == nil
}
