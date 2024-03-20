package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// time complexity is O(b + m), space complexity is O(1)
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	prevNode, aNode, bNode, list2LastNode := list1, list1, list1, list2
	// time complexity is O(b - a)
	for i := 0; i < b - a; i++ {
		bNode = bNode.Next
	}

	// time complexity is O(m)
	for list2LastNode.Next != nil {
		list2LastNode = list2LastNode.Next
	}

	// time complexity is O(a)
	for i := 0; i < a; i++ {
		prevNode = aNode
		aNode = aNode.Next
		bNode = bNode.Next
	}

	prevNode.Next = list2
	list2LastNode.Next = bNode.Next
	return list1
}
