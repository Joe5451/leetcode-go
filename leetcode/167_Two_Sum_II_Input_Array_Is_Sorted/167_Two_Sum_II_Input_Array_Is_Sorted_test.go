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
	numbers []int
	target int
}

type output struct {
	result []int
}

func Test_Problem167(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 167 ----------")

	questions := []question{
		{
			input{[]int{2, 7, 11, 15}, 9},
			output{[]int{1, 2}},
		},
		{
			input{[]int{2, 3, 4}, 6},
			output{[]int{1, 3}},
		},
		{
			input{[]int{-1, 0}, -1},
			output{[]int{1, 2}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := twoSum(input.numbers, input.target)
		fmt.Printf(", output: %v\n", output)
	}
}
