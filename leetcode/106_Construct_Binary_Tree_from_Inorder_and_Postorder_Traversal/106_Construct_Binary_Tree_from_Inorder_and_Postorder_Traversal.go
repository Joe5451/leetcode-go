package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// time complexity is O(n^2), space complexity is O(n)
func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(inorder) < 1 {
		return nil
	}

	rootVal := postorder[len(postorder) - 1]
	root := &TreeNode{rootVal, nil, nil}

	inorderRootIndex := -1
	// time complexity: O(n)
	for i, num := range inorder {
		if num == rootVal {
			inorderRootIndex = i
			break
		}
	}

	leftCount := len(inorder) - 1 - inorderRootIndex
	root.Left = buildTree(inorder[:inorderRootIndex], postorder[:len(postorder) - 1 - leftCount])
	root.Right = buildTree(inorder[inorderRootIndex + 1:], postorder[len(postorder) - 1 - leftCount:len(postorder) - 1])
	return root
}
