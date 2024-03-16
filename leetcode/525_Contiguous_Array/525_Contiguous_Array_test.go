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

func Test_Problem525(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 525 ----------")

	questions := []question{
		{
			input{[]int{0, 1}},
			output{2},
		},
		{
			input{[]int{0, 1, 0}},
			output{2},
		},
		{
			input{[]int{1, 0, 0, 0, 1, 0, 1}},
			output{4},
		},
		{
			input{[]int{0, 0, 0, 0, 1, 0, 1, 1, 0}},
			output{6},
		},
		{
			input{[]int{0, 0, 0, 1, 1, 1, 0}},
			output{6},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := findMaxLength(input.nums)
		fmt.Printf(", output: %v\n", output)
	}
}
