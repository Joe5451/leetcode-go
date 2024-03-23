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
	k int
}

type output struct {
	result int
}

func Test_Problem125(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 125 ----------")

	questions := []question{
		{
			input{[]int{3, 2, 1, 5, 6, 4}, 2},
			output{5},
		},
		{
			input{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4},
			output{4},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := findKthLargest(input.nums, input.k)
		fmt.Printf(", output: %v\n", output)
	}
}
