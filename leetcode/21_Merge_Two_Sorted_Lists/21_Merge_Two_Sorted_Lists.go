package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// time complexity is O(n), space complexity is O(1)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	headNode := &ListNode{}
	currentNode := headNode
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			currentNode.Next = list1
			list1 = list1.Next
		} else {
			currentNode.Next = list2
			list2 = list2.Next
		}

		currentNode = currentNode.Next
	}

	for list1 != nil {
		currentNode.Next = list1
		currentNode = currentNode.Next
		list1 = list1.Next
	}

	for list2 != nil {
		currentNode.Next = list2
		currentNode = currentNode.Next
		list2 = list2.Next
	}

	return headNode.Next
}
