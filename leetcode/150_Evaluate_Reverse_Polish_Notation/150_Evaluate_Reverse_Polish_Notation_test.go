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
	tokens []string
}

type output struct {
	value int
}

func Test_Problem150(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 150 ----------")

	questions := []question{
		{
			input{[]string{"2", "1", "+", "3", "*"}},
			output{9},
		},
		{
			input{[]string{"4", "13", "5", "/", "+"}},
			output{6},
		},
		{
			input{[]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}},
			output{22},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := evalRPN(input.tokens)
		fmt.Printf(", output: %v\n", output)
	}
}
