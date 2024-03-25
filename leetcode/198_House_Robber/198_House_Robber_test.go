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

func Test_Problem198(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 198 ----------")

	questions := []question{
		{
			input{[]int{1, 2, 3, 1}},
			output{4},
		},
		{
			input{[]int{2, 7, 9, 3, 1}},
			output{12},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := rob(input.nums)
		fmt.Printf(", output: %v\n", output)
	}
}
