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
func kthSmallest(root *TreeNode, k int) int {
	inorderVals := []int{}
	inorder(root, &inorderVals, k)
	val := inorderVals[k - 1]
	return val
}

func inorder(node *TreeNode, inorderVals *[]int, k int) {
	if node == nil || len(*inorderVals) >= k {
		return
	}

	inorder(node.Left, inorderVals, k)
	*inorderVals = append(*inorderVals, node.Val)
	inorder(node.Right, inorderVals, k)
}
