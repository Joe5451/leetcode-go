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
	points [][]int
}

type output struct {
	result int
}

func Test_Problem452(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 452 ----------")

	questions := []question{
		{
			input{[][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}}},
			output{2},
		},
		{
			input{[][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}},
			output{4},
		},
		{
			input{[][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}}},
			output{2},
		},
		{
			input{[][]int{{1, 2}}},
			output{1},
		},
		{
			input{[][]int{{3, 9}, {7, 12}, {3, 8}, {6, 8}, {9, 10}, {2, 9}, {0, 9}, {3, 9}, {0, 6}, {2, 8}}},
			output{2},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := findMinArrowShots(input.points)
		fmt.Printf(", output: %v\n", output)
	}
}
