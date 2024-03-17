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
}

type output struct {
	matrix [][]int
}

func Test_Problem48(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 48 ----------")

	questions := []question{
		{
			input{[][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}},
			output{[][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			}},
		},
		{
			input{[][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			}},
			output{[][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		rotate(input.matrix)
		fmt.Printf(", output: %v\n", input.matrix)
	}
}
