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
	strs []string
}

type output struct {
	output [][]string
}

func Test_Problem49(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 49 ----------")

	questions := []question{
		{
			input{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}},
			output{[][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		},
		{
			input{[]string{""}},
			output{[][]string{{""}}},
		},
		{
			input{[]string{"a"}},
			output{[][]string{{"a"}}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := groupAnagrams(input.strs)
		fmt.Printf(", output: %v\n", output)
	}
}
