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
	matrix [][]byte
}

type output struct {
	result int
}

func Test_Problem221(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 221 ----------")

	questions := []question{
		{
			input{[][]byte{
				{'1', '0', '1', '0', '0'},
				{'1', '0', '1', '1', '1'},
				{'1', '1', '1', '1', '1'},
				{'1', '0', '0', '1', '0'},
			}},
			output{4},
		},
		{
			input{[][]byte{
				{'0', '1'},
				{'1', '0'},
			}},
			output{1},
		},
		{
			input{[][]byte{
				{'0'},
			}},
			output{0},
		},
		{
			input{[][]byte{
				{'0'},
				{'1'},
			}},
			output{1},
		},
		{
			input{[][]byte{
				{'1', '0', '1', '0'},
				{'1', '0', '1', '1'},
				{'1', '0', '1', '1'},
				{'1', '1', '1', '1'},
			}},
			output{4},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := maximalSquare(input.matrix)
		fmt.Printf(", output: %v\n", output)
	}
}
