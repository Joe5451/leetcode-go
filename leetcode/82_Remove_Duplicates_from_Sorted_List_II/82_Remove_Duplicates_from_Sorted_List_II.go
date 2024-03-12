package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// time complexity is O(n), space complexity is O(1)
func deleteDuplicates(head *ListNode) *ListNode {
	var previous *ListNode
	current := head

	for current != nil && current.Next != nil {
		if current.Val == current.Next.Val {
			for current.Next != nil && current.Val == current.Next.Val {
				current.Next = current.Next.Next
			}

			current = current.Next
			if previous == nil {
				head = current
			} else {
				previous.Next = current
			}
			continue
		}

		previous = current
		current = current.Next
	}

	return head
}
