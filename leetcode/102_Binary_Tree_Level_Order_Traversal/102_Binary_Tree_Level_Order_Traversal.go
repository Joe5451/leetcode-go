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
func levelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		currentSize := len(queue)
		currentLevel := []int{}
		for i := 0; i < currentSize; i++ {
			if queue[i] == nil {
				continue
			}

			currentLevel = append(currentLevel, queue[i].Val)
			queue = append(queue, queue[i].Left)
			queue = append(queue, queue[i].Right)
		}
		if len(currentLevel) > 0 {
			result = append(result, currentLevel)
		}
		queue = queue[currentSize:]
	}

	return result
}
