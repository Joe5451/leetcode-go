package leetcode

import (
	"testing"
	"fmt"
)

type question struct {
	input
	output
}

type input struct {
	nums []int
	k int
}

type output struct {
	result int
}

func Test_Problem713(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 713 ----------")

	questions := []question{
		{
			input{[]int{10, 5, 2, 6}, 100},
			output{8},
		},
		{
			input{[]int{1, 2, 3}, 0},
			output{0},
		},
		{
			input{[]int{10, 5, 1, 19}, 100},
			output{9},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := numSubarrayProductLessThanK(input.nums, input.k)
		fmt.Printf(", output: %v\n", output)
	}
}
