package leetcode

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Neighbors []*Node
 * }
 */

// time complexity is O(V + E), space complexity is O(V)
// V is vertix , E is edge
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	visited := make(map[int]*Node)
	return dfs(node, visited)
}

func dfs(node *Node, visited map[int]*Node) *Node {
	if copyNode ,ok := visited[node.Val]; ok {
		return copyNode
	}

	copyNode := &Node{Val: node.Val}
	visited[copyNode.Val] = copyNode
	for _, neighborNode := range node.Neighbors {
		copyNode.Neighbors = append(copyNode.Neighbors, dfs(neighborNode, visited))
	}

	return copyNode
}
