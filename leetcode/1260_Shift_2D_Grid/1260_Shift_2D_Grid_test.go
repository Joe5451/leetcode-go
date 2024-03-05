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
	k int
}

type output struct {
	grid [][]int
}

func Test_validMountainArray(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 1260 ----------")

	questions := []question{
		{
			input{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 1},
			output{[][]int{{9, 1, 2}, {3, 4, 5}, {6, 7, 8}}},
		},
		{
			input{[][]int{{3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}, {12, 0, 21, 13}}, 4},
			output{[][]int{{12, 0, 21, 13}, {3, 8, 1, 9}, {19, 7, 2, 5}, {4, 6, 11, 10}}},
		},
		{
			input{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, 9},
			output{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}},
		},
	}

	fmt.Println("solution1:")

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := shiftGrid(input.grid, input.k)
		fmt.Printf(", output: %v\n", output)
	}

	fmt.Println("\nsolution2:")

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := shiftGrid2(input.grid, input.k)
		fmt.Printf(", output: %v\n", output)
	}
}
