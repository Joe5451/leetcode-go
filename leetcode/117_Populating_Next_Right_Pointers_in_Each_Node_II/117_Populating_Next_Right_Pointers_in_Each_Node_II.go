/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

// time complexity is O(n), space complexity is O(n)
func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	slice1 := []*Node{root}
	slice2 := []*Node{}

	for len(slice1) > 0 {
		nodes := []*Node{}
		for _, node := range slice1 {
			nodes = append(nodes, node)

			if node.Left != nil {
				slice2 = append(slice2, node.Left)
			}

			if node.Right != nil {
				slice2 = append(slice2, node.Right)
			}
		}

		for i := 0; i < len(nodes) - 1; i++ {
			nodes[i].Next = nodes[i + 1]
		}

		slice1 = slice2
		slice2 = []*Node{}
	}

	return root
}
