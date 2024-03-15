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

func Test_Problem162(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 162 ----------")

	questions := []question{
		{
			input{[]int{1, 2, 3, 1}},
			output{2},
		},
		{
			input{[]int{1, 2, 1, 3, 5, 6, 4}},
			output{5},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := findPeakElement(input.nums)
		fmt.Printf(", output: %v\n", output)
	}
}
