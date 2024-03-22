package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// time complexity is O(n), space complexity is O(n)
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	slow, fast := head, head
	forward := []int{}
	for fast != nil && fast.Next != nil {
		forward = append(forward, slow.Val)
		slow = slow.Next
		fast = fast.Next.Next
	}

	if fast != nil {
		slow = slow.Next
	}

	pos := len(forward) - 1
	for slow != nil {
		if slow.Val != forward[pos] {
			return false
		}

		pos--
		slow = slow.Next
	}

	return true
}
