package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// time complexity is O(n), space complexity is O(depth)
func smallestFromLeaf(root *TreeNode) string {
	ans := ""
	dfs(root, "", &ans)
	return ans
}

func dfs(root *TreeNode, currentString string, ans *string) {
	if root == nil {
		return
	}

	currentString = string(root.Val + 97) + currentString
	if root.Left == nil && root.Right == nil {
		if *ans == "" || currentString < *ans {
			*ans = currentString
		}
		return
	}

	dfs(root.Left, currentString, ans)
	dfs(root.Right, currentString, ans)
}
