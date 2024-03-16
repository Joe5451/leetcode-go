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
	matrix [][]int
	target int
}

type output struct {
	result bool
}

func Test_Problem74(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 74 ----------")

	questions := []question{
		{
			input{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3},
			output{true},
		},
		{
			input{[][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13},
			output{false},
		},
		{
			input{[][]int{{1}}, 0},
			output{false},
		},
		{
			input{[][]int{{1}, {3}}, 0},
			output{false},
		},
	}

	fmt.Println("solution1:")
	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := searchMatrix(input.matrix, input.target)
		fmt.Printf(", output: %v\n", output)
	}
}
