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
	order string
	s string
}

type output struct {
	result string
}

func Test_Problem791(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 791 ----------")

	questions := []question{
		{
			input{"cba", "abcd"},
			output{"cbad"},
		},
		{
			input{"bcafg", "abcd"},
			output{"bcad"},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := customSortString(input.order, input.s)
		fmt.Printf(", output: %v\n", output)
	}
}
