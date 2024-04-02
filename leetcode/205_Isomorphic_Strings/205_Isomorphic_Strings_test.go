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
	t string
}

type output struct {
	result bool
}

func Test_Problem205(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 205 ----------")

	questions := []question{
		{
			input{"egg", "add"},
			output{true},
		},
		{
			input{"foo", "bar"},
			output{false},
		},
		{
			input{"paper", "title"},
			output{true},
		},
		{
			input{"badc", "baba"},
			output{false},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := isIsomorphic(input.s, input.t)
		fmt.Printf(", output: %v\n", output)
	}
}
