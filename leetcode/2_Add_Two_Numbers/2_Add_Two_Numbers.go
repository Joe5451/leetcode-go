package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// time complexity is O(n), where n-> Maximum of lengths of l1 and l2
// space complexity is O(n), where n-> Maximum of lengths of l1 and l2
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	currentNode := head
	for l1 != nil || l2 != nil {
		if l1 != nil {
			currentNode.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			currentNode.Val += l2.Val
			l2 = l2.Next
		}

		if currentNode.Val >= 10 {
			currentNode.Val -= 10
			currentNode.Next = &ListNode{Val: 1}
		} else if l1 != nil || l2 != nil {
			currentNode.Next = &ListNode{}
		}

		currentNode = currentNode.Next
	}

	return head
}
