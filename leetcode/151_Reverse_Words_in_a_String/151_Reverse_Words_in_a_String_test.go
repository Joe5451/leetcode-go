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
}

type output struct {
	s string
}

func Test_Problem151(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 151 ----------")

	questions := []question{
		{
			input{"the sky is blue"},
			output{"blue is sky the"},
		},
		{
			input{"  hello world  "},
			output{"world hello"},
		},
		{
			input{"a good   example"},
			output{"example good a"},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := reverseWords(input.s)
		fmt.Printf(", output: %v\n", output)
	}
}
