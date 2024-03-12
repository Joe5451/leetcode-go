package leetcode

import "sort"

// time complexity is O(nlog(n)), space complexity is O(1)
// O(nlogn) for sort
func hIndex(citations []int) int {
	sort.Ints(citations)
	h := 0
	size := len(citations)
	for i, citation := range citations {
		if size - i >= citation && citation > h {
			h = citation
		} else if citation > size - i && size - i > h {
			h = size - i
		}
	}

	return h
}
