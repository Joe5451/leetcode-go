package leetcode

// time complexity is O(m+n), space complexity is O(n)
// m is length of order and would not exceed 26
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
