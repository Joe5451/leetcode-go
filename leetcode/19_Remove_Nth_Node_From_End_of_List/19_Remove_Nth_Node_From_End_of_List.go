package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//  time complexity is O(n), space complexity is O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	left, right := head, head

	for i := 0; i < n; i++ {
		right = right.Next
	}

	if right == nil {
		return head.Next
	}

	for right.Next != nil {
		right = right.Next
		left = left.Next
	}

	left.Next = left.Next.Next
	return head
}
