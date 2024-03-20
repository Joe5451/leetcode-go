package leetcode

// time complexity is O(V + E), space complexity is O(V + E)
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make([][]int, numCourses)
	for i := 0; i < len(prerequisites); i++ {
		graph[prerequisites[i][0]] = append(graph[prerequisites[i][0]], prerequisites[i][1])
	}

	visited := make([]int, numCourses)
	for i := 0; i < numCourses; i++ {
		if hasCycle(graph, visited, i) {
			return false
		}
	}

	return true
}

func hasCycle(graph [][]int, visited []int, node int) bool {
	if visited[node] == 1 {
		return true
	}

	if visited[node] == -1 {
		return false
	}

	visited[node] = 1

	for _, neighbor := range graph[node] {
		if hasCycle(graph, visited, neighbor) {
			return true
		}
	}

	visited[node] = -1 // mark no cycle of the node
	return false
}
