package leetcode

// time complexity is O(n), space complexity is O(n)
func checkValidString(s string) bool {
	starIdx := []int{}
	leftParenthesisIdx := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			starIdx = append(starIdx, i)
		} else if s[i] == ')' {
			if len(leftParenthesisIdx) > 0 {
				leftParenthesisIdx = leftParenthesisIdx[:len(leftParenthesisIdx) - 1]
			} else if len(starIdx) > 0 {
				starIdx = starIdx[:len(starIdx) - 1]
			} else {
				return false
			}
		} else {
			leftParenthesisIdx = append(leftParenthesisIdx, i)
		}
	}

	starPos := 0
	for _, i := range leftParenthesisIdx {
		for starPos < len(starIdx) && starIdx[starPos] < i {
			starPos++
		}

		if starPos >= len(starIdx) {
			return false
		} else {
			starPos++
		}
	}

	return true
}
