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
	target int
	nums []int
}

type output struct {
	result int
}

func Test_Problem209(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 209 ----------")

	questions := []question{
		{
			input{7, []int{2, 3, 1, 2, 4, 3}},
			output{2},
		},
		{
			input{4, []int{1, 4, 4}},
			output{1},
		},
		{
			input{11, []int{1, 1, 1, 1, 1, 1, 1, 1}},
			output{0},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := minSubArrayLen(input.target, input.nums)
		fmt.Printf(", output: %v\n", output)
	}
}
