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
	result int
}

func Test_Problem41(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 41 ----------")

	questions := []question{
		{
			input{[]int{1, 2, 0}},
			output{3},
		},
		{
			input{[]int{3, 4, -1, 1}},
			output{2},
		},
		{
			input{[]int{7, 8, 9, 11, 12}},
			output{1},
		},
		{
			input{[]int{1, 1}},
			output{2},
		},
		{
			input{[]int{2, 1}},
			output{3},
		},
		{
			input{[]int{1, 2, 6, 3, 5, 4}},
			output{7},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := firstMissingPositive(input.nums)
		fmt.Printf(", output: %v\n", output)
	}
}
