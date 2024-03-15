package leetcode

import "strings"

// time complexity is O(n), space complexity is O(1)
func reverseWords(s string) string {
	result := ""
	word := ""

	for _, char := range s {
		if string(char) != " " {
			word += string(char)
		} else if word != "" {
			if result != "" {
				result = word + " " + result
			} else {
				result = word
			}

			word = ""
		}
	}

	if word != "" {
		if result != "" {
			result = word + " " + result
		} else {
			result = word
		}
	}

	return result
}

func reverseWords2(s string) string {
	words := strings.Fields(s)
	result := ""
	for i := len(words) - 1; i >= 0; i-- {
		result += words[i]

		if i != 0 {
			result += " "
		}
	}

	return result
}
