package leetcode

// time complexity is O(m+n), space complexity is O(1)
func getCommon(nums1 []int, nums2 []int) int {
	for i, j := 0, 0; i < len(nums1) && j < len(nums2); {
		if nums1[i] == nums2[j] {
			return nums1[i]
		}
		if nums1[i] > nums2[j] {
			j++
		} else {
			i++
		}
	}

	return -1
}
