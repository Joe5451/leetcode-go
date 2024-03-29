package leetcode

// time complexity is O(n), space complexity is O(1)
func countSubarrays(nums []int, k int) int64 {
	maxNum := nums[0]
	maxNumCount := 0
	maxNumStartIndex := 0
	for i, num := range nums {
		if num > maxNum {
			maxNum = num
			maxNumCount = 1
			maxNumStartIndex = i
		} else if num == maxNum {
			maxNumCount++
		}
	}

	if maxNumCount < k {
		return 0
	}

	result := 0
	currentCount := 0
	for start, end := maxNumStartIndex, maxNumStartIndex; end < len(nums); end++ {
		if nums[end] == maxNum {
			currentCount++

			for currentCount > k {
				if nums[start] == maxNum {
					currentCount--
					start++
				}
			}

			for nums[start] != maxNum {
				start++
			}

			if currentCount == k {
				result += start + 1

				left := end + 1
				for left < len(nums) && nums[left] != maxNum {
					result += start + 1
					left++
				}
			}
		}
	}

	return int64(result)
}
