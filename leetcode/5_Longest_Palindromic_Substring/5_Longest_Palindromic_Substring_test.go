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
	result string
}

func Test_Problem5(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 5 ----------")

	questions := []question{
		{
			input{"babad"},
			output{"bab"},
		},
		{
			input{"cbbd"},
			output{"bb"},
		},
		{
			input{"aaaa"},
			output{"aaaa"},
		},
		{
			input{"aaaaa"},
			output{"aaaaa"},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := longestPalindrome(input.s)
		fmt.Printf(", output: %v\n", output)
	}
}
