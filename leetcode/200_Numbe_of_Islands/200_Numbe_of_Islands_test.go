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
	grid [][]byte
}

type output struct {
	result int
}

func Test_Problem200(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 200 ----------")

	questions := []question{
		{
			input{[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			}},
			output{1},
		},
		{
			input{[][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			}},
			output{3},
		},
		{
			input{[][]byte{
				{'1', '1', '1'},
				{'0', '1', '0'},
				{'1', '1', '1'},
			}},
			output{1},
		},

	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := numIslands(input.grid)
		fmt.Printf(", output: %v\n", output)
	}
}
