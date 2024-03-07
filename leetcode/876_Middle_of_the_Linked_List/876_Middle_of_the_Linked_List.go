package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// time complexity is O(n), space complexity is O(1)
func middleNode(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	left := head
	right := head.Next

	for right != nil && right.Next != nil {
		left = left.Next
		right = right.Next.Next
	}

	if right == nil {
		return left
	}

	return left.Next
}
