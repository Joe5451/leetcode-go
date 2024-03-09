package leetcode

// time complexity is O(n), space complexity is O(n)

/**
 * ranging map[int]bool in some case is faster than ranging []int.
 * for example: var nums = []int{0, 1, 1, 1, 2}
 * map just store the value as the key once
 * so ranging map iteration count is less than ranging slice
 */

func longestConsecutive(nums []int) int {
	numsSet := map[int]bool{}
	for _, num := range nums {
		numsSet[num] = true
	}

	maxSequence := 0
	// for _, num := range nums {
	for num, _ := range numsSet {
		if numsSet[num - 1] {
			continue
		}

		currentNum, sequence := num + 1, 1
		for numsSet[num + sequence] {
			currentNum++
			sequence++
		}

		if sequence > maxSequence {
			maxSequence = sequence
		}
	}

	return maxSequence
}
