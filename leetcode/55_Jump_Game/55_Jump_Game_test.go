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
}

type output struct {
	result bool
}

func Test_Problem55(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 55 ----------")

	questions := []question{
		{
			input{[]int{2, 3, 1, 1, 4}},
			output{true},
		},
		{
			input{[]int{3, 2, 1, 0, 4}},
			output{false},
		},
	}

	for _, question := range questions {
		nums := question.input.nums
		fmt.Printf("input: %v", question.input)
		output := canJump(nums)
		fmt.Printf(", output: %v\n", output)
	}
}
