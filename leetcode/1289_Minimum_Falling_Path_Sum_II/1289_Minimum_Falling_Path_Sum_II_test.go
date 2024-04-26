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

func Test_Problem1289(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 1289 ----------")

	questions := []question{
		{
			input{[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}},
			output{13},
		},
		{
			input{[][]int{
				{7},
			}},
			output{7},
		},
		{
			input{[][]int{
				{-37, 51, -36, 34, -22},
				{82, 4, 30, 14, 38},
				{-68, -52, -92, 65, -85},
				{-49, -3, -77, 8, -19},
				{-60, -71, -21, -62, -73},
			}},
			output{-268},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := minFallingPathSum(input.grid)
		fmt.Printf(", output: %v\n", output)
	}
}
