package leetcode

// time complexity is O(n), space complexity is O(n)
func groupAnagrams(strs []string) [][]string {
	letterFrequencySet := make(map[[26]int][]string)

	for _, str := range strs {
		letterFrequency := getLetterFrequency(str)
		letterFrequencySet[letterFrequency] = append(letterFrequencySet[letterFrequency], str)
	}

	anagrams := make([][]string, len(letterFrequencySet))
	pos := 0
	for _, anagram := range letterFrequencySet {
		anagrams[pos] = anagram
		pos++
	}

	return anagrams
}

// hash
func getLetterFrequency(word string) [26]int {
	res := [26]int{}
	for _, char := range word {
		res[char - 'a']++
	}
	return res
}
