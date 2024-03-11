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
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	if len(preorder) == 1 {
		return &TreeNode{preorder[0], nil, nil}
	}

	root := TreeNode{preorder[0], nil, nil}

	var inorderRootIndex int
	for i, val := range inorder {
		if val == preorder[0] {
			inorderRootIndex = i
			break
		}
	}

	root.Left = buildTree(preorder[1:inorderRootIndex + 1], inorder[:inorderRootIndex])
	root.Right = buildTree(preorder[inorderRootIndex + 1:], inorder[inorderRootIndex + 1:])
	return &root
}
