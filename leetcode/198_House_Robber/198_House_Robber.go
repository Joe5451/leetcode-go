package leetcode

// time complexity is O(n), space complexity is O(n)
func rob(nums []int) int {
	house := 0
	visited := make(map[int]int)
	return dfs(house, nums, visited)
}

func dfs(house int, nums []int, visited map[int]int) int {
	if money, ok := visited[house]; ok {
		return money
	}

	robMoney := 0
	if house + 2 < len(nums) {
		robMoney = max(nums[house] + dfs(house + 2, nums, visited), dfs(house + 1, nums, visited))
	} else if house + 1 < len(nums) {
		robMoney = max(nums[house], nums[house + 1])
	} else if house < len(nums) {
		robMoney = nums[house]
	}

	visited[house] = robMoney
	return robMoney
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}

	return num2
}
