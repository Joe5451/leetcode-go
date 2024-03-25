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
	wordDict []string
}

type output struct {
	result bool
}

func Test_Problem139(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 139 ----------")

	questions := []question{
		{
			input{"leetcode", []string{"leet", "code"}},
			output{true},
		},
		{
			input{"applepenapple", []string{"apple", "pen"}},
			output{true},
		},
		{
			input{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}},
			output{false},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := wordBreak(input.s, input.wordDict)
		fmt.Printf(", output: %v\n", output)
	}
}
