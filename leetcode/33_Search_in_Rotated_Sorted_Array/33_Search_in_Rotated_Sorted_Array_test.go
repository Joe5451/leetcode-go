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
	target int
}

type output struct {
	result int
}

func Test_Problem33(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 33 ----------")

	questions := []question{
		{
			input{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
			output{4},
		},
		{
			input{[]int{4, 5, 6, 7, 0, 1, 2}, 3},
			output{-1},
		},
		{
			input{[]int{1}, 0},
			output{-1},
		},
		{
			input{[]int{1}, 1},
			output{0},
		},
		{
			input{[]int{3, 1}, 1},
			output{1},
		},
		{
			input{[]int{4, 5, 6, 7, 0, 1, 2}, 5},
			output{1},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := search(input.nums, input.target)
		fmt.Printf(", output: %v\n", output)
	}
}
