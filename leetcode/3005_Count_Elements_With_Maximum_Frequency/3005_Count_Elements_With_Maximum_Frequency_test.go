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
	count int
}

func Test_Problem3005(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 3005 ----------")

	questions := []question{
		{
			input{[]int{1, 2, 2, 3, 1, 4}},
			output{4},
		},
		{
			input{[]int{1, 2, 3, 4, 5}},
			output{5},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := maxFrequencyElements(input.nums)
		fmt.Printf(", output: %v\n", output)
	}
}
