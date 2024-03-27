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
	grid [][]int
}

type output struct {
	result int
}

func Test_Problem64(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 64 ----------")

	questions := []question{
		{
			input{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}},
			output{7},
		},
		{
			input{[][]int{{1, 2, 3}, {4, 5, 6}}},
			output{12},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := minPathSum(input.grid)
		fmt.Printf(", output: %v\n", output)
	}
}
