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

func Test_Problem1137(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 1137 ----------")

	questions := []question{
		{
			input{4},
			output{4},
		},
		{
			input{25},
			output{1389537},
		},
		{
			input{35},
			output{615693474},
		},
		{
			input{0},
			output{0},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := tribonacci(input.n)
		fmt.Printf(", output: %v\n", output)
	}
}
