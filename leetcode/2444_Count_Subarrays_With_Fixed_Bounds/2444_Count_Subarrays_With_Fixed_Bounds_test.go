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
	minK int
	maxK int
}

type output struct {
	result int
}

func Test_Problem2444(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 2444 ----------")

	questions := []question{
		{
			input{[]int{1, 3, 5, 2, 7, 5}, 1, 5},
			output{2},
		},
		{
			input{[]int{1, 1, 1, 1}, 1, 1},
			output{10},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := countSubarrays(input.nums, input.minK, input.maxK)
		fmt.Printf(", output: %v\n", output)
	}
}
