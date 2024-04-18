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

func Test_Problem463(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 463 ----------")

	questions := []question{
		{
			input{[][]int{
				{0, 1, 0, 0},
				{1, 1, 1, 0},
				{0, 1, 0, 0},
				{1, 1, 0, 0},
			}},
			output{16},
		},
		{
			input{[][]int{
				{1},
			}},
			output{4},
		},
		{
			input{[][]int{
				{1},
				{0},
			}},
			output{4},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := islandPerimeter(input.grid)
		fmt.Printf(", output: %v\n", output)
	}
}
