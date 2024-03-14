package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// time complexity is O(n), space complexity is O(1)
func rotateRight(head *ListNode, k int) *ListNode {
	listLength := 0
	for current := head; current != nil; {
		listLength++
		current = current.Next
	}

	if listLength == 0 || k % listLength == 0 {
		return head
	}

	k = k % listLength
	slowNode := head
	fastNode := head
	for i := 0; i < k; i++ {
		fastNode = fastNode.Next
	}
	for fastNode.Next != nil {
		slowNode = slowNode.Next
		fastNode = fastNode.Next
	}

	newHead := slowNode.Next
	slowNode.Next = nil
	fastNode.Next = head
	return newHead
}
