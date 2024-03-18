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
	board [][]byte
}

type output struct {
	board [][]byte
}

func Test_Problem130(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 130 ----------")

	questions := []question{
		{
			input{[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'X', 'O', 'X'},
				{'X', 'O', 'X', 'X'},
			}},
			output{[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'X', 'X'},
			}},
		},
		{
			input{[][]byte{
				{'X'},
			}},
			output{[][]byte{
				{'X'},
			}},
		},
		{
			input{[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'O', 'O', 'X'},
				{'O', 'O', 'X', 'X'},
			}},
			output{[][]byte{
				{'X', 'X', 'X', 'X'},
				{'X', 'O', 'O', 'X'},
				{'X', 'O', 'O', 'X'},
				{'O', 'O', 'X', 'X'},
			}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		solve(input.board)
		fmt.Printf(", output: %v\n", input.board)
	}
}
