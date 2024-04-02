package leetcode

// time complexity is O(n), space complexity is O(n)
func isIsomorphic(s string, t string) bool {
	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if result, ok := sMap[s[i]]; ok {
			if result != t[i] || s[i] != tMap[t[i]] {
				return false
			}
		} else {
			if _, ok := tMap[t[i]]; ok {
				return false
			}
			sMap[s[i]] = t[i]
			tMap[t[i]] = s[i]
		}
	}
	return true
}

// time complexity is O(n), space complexity is O(n)
func isIsomorphic2(s string, t string) bool {
	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)
	for i := 0; i < len(s); i++ {
		if result, ok := sMap[s[i]]; ok && result != t[i] {
			return false
		}
		sMap[s[i]] = t[i]

		if result, ok := tMap[t[i]]; ok && result != s[i] {
			return false
		}
		tMap[t[i]] = s[i]
	}
	return true
}
