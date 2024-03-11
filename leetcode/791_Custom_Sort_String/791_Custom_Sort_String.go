package leetcode

// time complexity is O(n), space complexity is O(n)
func customSortString(order string, s string) string {
	result := ""
	sCount := make(map[rune]int)
	for _, str := range s {
		sCount[str]++
	}

	for _, str := range order {
		if count, ok := sCount[str]; ok {
			for i := 0; i < count; i++ {
				result += string(str)
			}

			delete(sCount, str)
		}
	}

	for str, count := range sCount {
		for i := 0; i < count; i++ {
			result += string(str)
		}
	}

	return result
}
