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
	height []int
}

type output struct {
	result int
}

func Test_Problem11(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 11 ----------")

	questions := []question{
		{
			input{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			output{49},
		},
		{
			input{[]int{1, 1}},
			output{1},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := maxArea(input.height)
		fmt.Printf(", output: %v\n", output)
	}
}
