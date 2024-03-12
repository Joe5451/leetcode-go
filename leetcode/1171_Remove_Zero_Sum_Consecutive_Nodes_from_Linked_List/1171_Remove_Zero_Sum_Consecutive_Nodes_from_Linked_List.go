package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// time complexity is O(n), space complexity is O(n)
func removeZeroSumSublists(head *ListNode) *ListNode {
	sumToNode := make(map[int]*ListNode)
	dummy := &ListNode{0, head}
	current := dummy.Next
	sum := 0
	sumToNode[0] = dummy
	for current != nil {
		sum += current.Val

		if node, ok := sumToNode[sum]; ok {
			deleteSum := sum
			deleteNode := node.Next
			// delete zero sum consecutive nodes map
			for deleteNode != current {
				deleteSum += deleteNode.Val
				// remain dummy map
				if deleteSum != 0 {
					delete(sumToNode, deleteSum)
				}
				deleteNode = deleteNode.Next
			}
			node.Next = current.Next
		} else {
			sumToNode[sum] = current
		}

		current = current.Next
	}

	return dummy.Next
}
