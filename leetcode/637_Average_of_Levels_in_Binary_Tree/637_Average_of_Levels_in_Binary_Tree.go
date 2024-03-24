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
func averageOfLevels(root *TreeNode) []float64 {
	result := []float64{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		currentSize := len(queue)
		sum := 0
		for i := 0; i < currentSize; i++ {
			sum += queue[i].Val
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}

		result = append(result, float64(sum) / float64(currentSize))
		queue = queue[currentSize:]
	}

	return result
}
