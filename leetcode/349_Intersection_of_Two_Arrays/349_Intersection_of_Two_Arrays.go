package leetcode

// time complexity is O(m+n), space complexity is O(m+n)
func intersection(nums1 []int, nums2 []int) []int {
	nums1Set := make(map[int]bool)
	for _, num := range nums1 {
		nums1Set[num] = true
	}

	intersectedSet := make(map[int]bool)
	for _, num := range nums2 {
		if nums1Set[num] {
			intersectedSet[num] = true
		}
	}

	result := []int{}
	for num := range intersectedSet {
		result = append(result, num)
	}

	return result
}
