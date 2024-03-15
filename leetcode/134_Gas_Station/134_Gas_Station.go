package leetcode

// time complexity is O(n), space complexity is O(1)
func canCompleteCircuit(gas []int, cost []int) int {
	positive := 0
	negative := 0
	start := -1
	for i := 0; i < len(gas); i++ {
		positive = positive + gas[i] - cost[i]
		if positive < 0 {
			negative += positive
			positive = 0
			start = -1
		} else if start == -1 {
			start = i
		}
	}

	if positive + negative >= 0 {
		return start
	}

	return -1
}
