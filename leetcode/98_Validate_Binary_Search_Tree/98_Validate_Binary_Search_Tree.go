package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// time complexity is O(n)
// space complexity is O(h), where h is tree height.
// For a balanced tree, h is log(n), making the space complexity O(log(n)).
// In the worst case of a skewed tree, h is n, making the space complexity O(n).
func isValidBST(root *TreeNode) bool {
	return dfs(root, nil, nil)
}

func dfs(root *TreeNode, min, max *TreeNode) bool {
	if root == nil {
		return true
	}

	if (max != nil && root.Val >= max.Val) || (min != nil && root.Val <= min.Val) {
		return false
	}

	return dfs(root.Left, min, root) && dfs(root.Right, root, max)
}
