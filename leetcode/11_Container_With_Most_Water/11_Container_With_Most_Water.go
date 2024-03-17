package leetcode

func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxArea := 0
	for left < right {
		minHeight := min(height[left], height[right])
		area := (right - left) * minHeight
		if area > maxArea {
			maxArea = area
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func min(height1 int, height2 int) int {
	if height1 > height2 {
		return height2
	}

	return height1
}
