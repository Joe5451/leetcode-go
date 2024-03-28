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

func Test_Problem2958(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 2958 ----------")

	questions := []question{
		{
			input{[]int{1, 2, 3, 1, 2, 3, 1, 2}, 2},
			output{6},
		},
		{
			input{[]int{1, 2, 1, 2, 1, 2, 1, 2}, 1},
			output{2},
		},
		{
			input{[]int{5, 5, 5, 5, 5, 5, 5}, 4},
			output{4},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := maxSubarrayLength(input.nums, input.k)
		fmt.Printf(", output: %v\n", output)
	}
}
