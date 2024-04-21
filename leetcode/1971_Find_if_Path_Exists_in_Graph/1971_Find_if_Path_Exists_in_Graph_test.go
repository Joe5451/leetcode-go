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
	n int
	edges [][]int
	source int
	destination int
}

type output struct {
	result bool
}

func Test_Problem1971(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 1971 ----------")

	questions := []question{
		{
			input{
				3,
				[][]int{{0, 1}, {1, 2}, {2, 0}},
				0,
				2,
			},
			output{true},
		},
		{
			input{
				6,
				[][]int{{0, 1}, {0, 2}, {3, 5}, {5, 4}, {4, 3}},
				0,
				5,
			},
			output{false},
		},
		{
			input{
				1,
				[][]int{},
				0,
				0,
			},
			output{true},
		},
		{
			input{
				10,
				[][]int{{4, 3}, {1, 4}, {4, 8}, {1, 7}, {6, 4}, {4, 2}, {7, 4}, {4, 0}, {0, 9}, {5, 4}},
				5,
				9,
			},
			output{true},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := validPath(input.n, input.edges, input.source, input.destination)
		fmt.Printf(", output: %v\n", output)
	}
}
