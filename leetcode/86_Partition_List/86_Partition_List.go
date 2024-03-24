package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// time complexity is O(n), space complexity is O(1)
func partition(head *ListNode, x int) *ListNode {
	smallDummy, bigDummy := &ListNode{0, nil}, &ListNode{0, nil}
	small, big, current := smallDummy, bigDummy, head
	for current != nil {
		if current.Val < x {
			small.Next = current
			small = current
		} else {
			big.Next = current
			big = current
		}
		current = current.Next
	}

	small.Next = bigDummy.Next
	big.Next = nil
	return smallDummy.Next
}
