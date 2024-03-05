package leetcode

// time complexity is O(n), space complexity is O(1)
func validMountainArray(arr []int) bool {
	i := 1
	for i < len(arr) && arr[i] > arr[i - 1] {
		i++
	}

	if i == 1 || i >= len(arr) {
		return false
	}

	for i < len(arr) && arr[i] < arr[i - 1] {
		i++
	}

	if i != len(arr) {
		return false
	}

	return true
}

// time complexity is O(n), space complexity is O(1)
func validMountainArray2(arr []int) bool {
	isUp, isDown := false, false

	for i := 1; i < len(arr); i++ {
		if !isDown && arr[i] > arr[i - 1] {
			isUp = true
			continue
		}

		if isUp && arr[i] < arr[i - 1] {
			isDown = true
			continue
		}

		return false
	}

	return isUp && isDown
}
