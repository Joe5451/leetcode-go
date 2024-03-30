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
	s1 string
	s2 string
	s3 string
}

type output struct {
	result bool
}

func Test_Problem97(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 97 ----------")

	questions := []question{
		{
			input{"aabcc", "dbbca", "aadbbcbcac"},
			output{true},
		},
		{
			input{"aabcc", "dbbca", "aadbbbaccc"},
			output{false},
		},
		{
			input{"", "", ""},
			output{true},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := isInterleave(input.s1, input.s2, input.s3)
		fmt.Printf(", output: %v\n", output)
	}
}
