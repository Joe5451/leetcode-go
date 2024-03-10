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
	result int
}

func Test_Problem3(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 3 ----------")

	questions := []question{
		{
			input{"abcabcbb"},
			output{3},
		},
		{
			input{"bbbbb"},
			output{1},
		},
		{
			input{"pwwkew"},
			output{3},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := lengthOfLongestSubstring(input.s)
		fmt.Printf(", output: %v\n", output)
	}
}
