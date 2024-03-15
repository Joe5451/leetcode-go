package leetcode

// time complexity is O(n), space complexity is O(n)
func productExceptSelf(nums []int) []int {
	products := make([]int, len(nums))
	product := 1
	// reverse product
	for i := len(nums) - 1; i >= 0; i-- {
		product = product*nums[i]
		products[i] = product
	}

	product = 1
	for i, num := range nums {
		if i + 1 < len(products) {
			products[i] = product * products[i + 1]
		} else {
			products[i] = product
		}

		product = product*num
	}

	return products
}

func productExceptSelf2(nums []int) []int {
	products := make([]int, len(nums))
	product := 1
	for i, num := range nums {
		products[i] = product
		product *= num
	}

	product = 1
	for i := len(nums) - 1; i >= 0; i-- {
		products[i] *= product
		product *= nums[i]
	}

	return products
}
