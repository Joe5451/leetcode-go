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
	s string
	numRows int
}

type output struct {
	s string
}

func Test_Problem6(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 6 ----------")

	questions := []question{
		{
			input{"PAYPALISHIRING", 3},
			output{"PAHNAPLSIIGYIR"},
		},
		{
			input{"PAYPALISHIRING", 4},
			output{"PINALSIGYAHRPI"},
		},
		{
			input{"A", 1},
			output{"A"},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := convert(input.s, input.numRows)
		fmt.Printf(", output: %v\n", output)
	}
}
