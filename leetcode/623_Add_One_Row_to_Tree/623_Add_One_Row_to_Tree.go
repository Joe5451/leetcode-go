package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// BFS
// time complexity is O(n)
// space complexity is O(x) where x is max number of tree nodes of the target level
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	dummy := &TreeNode{Val: 0, Left: root}
	currentRow := []*TreeNode{dummy}

	for curDepth := 0; curDepth < depth - 1; curDepth++ {
		currentLen := len(currentRow)
		for i := 0; i < currentLen; i++ {
			if currentRow[i].Left != nil {
				currentRow = append(currentRow, currentRow[i].Left)
			}
			if currentRow[i].Right != nil {
				currentRow = append(currentRow, currentRow[i].Right)
			}
		}
		currentRow = currentRow[currentLen:]
	}

	for _, node := range currentRow {
		newLeftNode := &TreeNode{Val: val, Left: node.Left}
		newRightNode := &TreeNode{Val: val, Right: node.Right}
		node.Left, node.Right = newLeftNode, newRightNode
	}
	return dummy.Left
}
