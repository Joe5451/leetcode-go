package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// time complexity is O(n), space complexity is O(1)
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	current := head
	next := head.Next
	current.Next = nil
	for next != nil {
		temp := next.Next
		next.Next = current
		current = next
		next = temp
	}

	return current
}
