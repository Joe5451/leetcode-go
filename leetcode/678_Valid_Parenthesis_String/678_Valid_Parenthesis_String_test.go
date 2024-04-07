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
	result bool
}

func Test_Problem678(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 678 ----------")

	questions := []question{
		{
			input{"()"},
			output{true},
		},
		{
			input{"(*)"},
			output{true},
		},
		{
			input{"(*))"},
			output{true},
		},
		{
			input{"*()"},
			output{true},
		},
		{
			input{"(((((()*)(*)*))())())(()())())))((**)))))(()())()"},
			output{false},
		},
		{
			input{"(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())"},
			output{false},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := checkValidString(input.s)
		fmt.Printf(", output: %v\n", output)
	}
}
