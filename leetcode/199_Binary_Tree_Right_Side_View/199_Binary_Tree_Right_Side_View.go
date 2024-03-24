package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// time complexity is O(n), space complexity is O(n)
func rightSideView(root *TreeNode) []int {
	result := []int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		currentSize := len(queue)
		rightSideVal := -101
		for i := 0; i < currentSize; i++ {
			currentNode := queue[i]
			if currentNode == nil {
				continue
			}

			rightSideVal = currentNode.Val
			queue = append(queue, currentNode.Left)
			queue = append(queue, currentNode.Right)
		}
		if rightSideVal != -101 {
			result = append(result, rightSideVal)
		}
		queue = queue[currentSize:]
	}

	return result
}
