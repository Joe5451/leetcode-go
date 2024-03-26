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
	n int
}

type output struct {
	result int
}

func Test_Problem70(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 70 ----------")

	questions := []question{
		{
			input{2},
			output{2},
		},
		{
			input{3},
			output{3},
		},
		{
			input{44},
			output{1134903170},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := climbStairs(input.n)
		fmt.Printf(", output: %v\n", output)
	}
}
