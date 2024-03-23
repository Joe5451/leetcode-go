package leetcode

// time complexity is O(n^2), space complexity is O(n^2)
func minMutation(startGene string, endGene string, bank []string) int {
	count := 0
	queue := []string{startGene}
	visited := make(map[string]bool)
	graph := make(map[string][]string)

	// init graph, time complexity: O(n^2), space complexity: O(n^2)
	graph[startGene] = getOneDiffGenes(startGene, bank)
	for _, bankGene := range bank {
		if bankGene == endGene {
			continue
		}

		graph[bankGene] = getOneDiffGenes(bankGene, bank)
	}

	// BFS, time complexity: O(n^2)
	for len(queue) > 0 {
		currentSize := len(queue)
		for i := 0; i < currentSize; i++ {
			currentGene := queue[i]
			if currentGene == endGene {
				return count
			}

			visited[currentGene] = true

			if genes, ok := graph[currentGene]; ok {
				for _, gene := range genes {
					if !visited[gene] {
						queue = append(queue, gene)
					}
				}
			}
		}

		queue = queue[currentSize:]
		count++
	}

	return -1
}

// time complexity: O(k*n) -> O(n), k is gene length
func getOneDiffGenes(gene string, bank []string) []string {
	result := []string{}
	for _, bankGene := range bank {
		diffCount := 0
		for i := 0; i < len(gene); i++ {
			if gene[i] != bankGene[i] {
				diffCount++
			}

			if diffCount > 1 {
				break
			}
		}

		if diffCount == 1 {
			result = append(result, bankGene)
		}
	}

	return result
}
