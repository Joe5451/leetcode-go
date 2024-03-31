package leetcode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

// time complexity is O(n), space complexity is O(n)
func copyRandomList(head *Node) *Node {
	orderToNew := make(map[*Node]*Node)
	current := head
	for current != nil {
		orderToNew[current] = &Node{current.Val, nil, nil}
		current = current.Next
	}

	current = head
	for current != nil {
		orderToNew[current].Next = orderToNew[current.Next]
		orderToNew[current].Random = orderToNew[current.Random]
		current = current.Next
	}

	return orderToNew[head]
}
