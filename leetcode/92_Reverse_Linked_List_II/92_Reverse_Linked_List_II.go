package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// this solution is better
// time complexity is O(n), space complexity is O(1)
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummy := &ListNode{0, head}
	prev := dummy
	for i := 0; i < left - 1; i++ {
		prev = prev.Next
	}

	current := prev.Next
	for i := 0; i < right - left; i++ {
		next := current.Next
		current.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return dummy.Next
}

// time complexity is O(n), space complexity is O(1)
func reverseBetween2(head *ListNode, left int, right int) *ListNode {
	if left == right {
		return head
	}
	var leftNode, rightNode, startNode, endNode, previousNode *ListNode
	currentNode := head
	leftNode = head

	count := 1
	for currentNode != nil {
		if count > right {
			break
		}

		if count == left - 1 {
			startNode = currentNode
			leftNode = currentNode.Next
		}
		if count == right {
			endNode = currentNode.Next
			rightNode = currentNode
		}

		if count >= left && count <= right {
			nextNode := currentNode.Next
			currentNode.Next = previousNode
			previousNode = currentNode
			currentNode = nextNode
			count++
			continue
		}

		count++
		previousNode = currentNode
		currentNode = currentNode.Next
	}

	leftNode.Next = endNode
	if startNode != nil {
		startNode.Next = rightNode
		return head
	} else {
		return rightNode
	}
}
