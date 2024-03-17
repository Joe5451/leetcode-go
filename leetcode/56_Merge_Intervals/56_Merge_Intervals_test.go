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
	intervals [][]int
}

type output struct {
	result [][]int
}

func Test_Problem56(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 56 ----------")

	questions := []question{
		{
			input{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}},
			output{[][]int{{1, 6}, {8, 10}, {15, 18}}},
		},
		{
			input{[][]int{{1, 4}, {4, 5}}},
			output{[][]int{{1, 5}}},
		},
		{
			input{[][]int{{1, 4}, {0, 0}}},
			output{[][]int{{0, 0}, {1, 4}}},
		},
		{
			input{[][]int{{1, 4}, {0, 2}, {3, 5}}},
			output{[][]int{{0, 5}}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := merge(input.intervals)
		fmt.Printf(", output: %v\n", output)
	}
}
