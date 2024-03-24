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
func zigzagLevelOrder(root *TreeNode) [][]int {
	result := [][]int{}
	queue := []*TreeNode{root}
	isRightFirst := true

	for len(queue) > 0 {
		currentSize := len(queue)
		currentLevel := []int{}
		tempQueue := []*TreeNode{}
		for i := 0; i < currentSize; i++ {
			if queue[i] == nil {
				continue
			}

			currentLevel = append(currentLevel, queue[i].Val)
			if isRightFirst {
				tempQueue = append([]*TreeNode{queue[i].Right, queue[i].Left}, tempQueue...)
			} else {
				tempQueue = append([]*TreeNode{queue[i].Left, queue[i].Right}, tempQueue...)
			}
		}

		if len(currentLevel) > 0 {
			result = append(result, currentLevel)
		}

		isRightFirst = !isRightFirst
		queue = tempQueue
	}

	return result
}
