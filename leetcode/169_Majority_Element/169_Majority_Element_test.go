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
	nums []int
}

type output struct {
	result int
}

func Test_Problem169(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 169 ----------")

	questions := []question{
		{
			input{[]int{3, 2, 3}},
			output{3},
		},
		{
			input{[]int{2, 2, 1, 1, 1, 2, 2}},
			output{2},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := majorityElement(input.nums)
		fmt.Printf(", output: %v\n", output)
	}
}
