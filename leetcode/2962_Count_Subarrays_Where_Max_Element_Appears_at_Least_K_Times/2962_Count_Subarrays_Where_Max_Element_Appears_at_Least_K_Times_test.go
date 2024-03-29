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

func Test_Problem2962(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 2962 ----------")

	questions := []question{
		{
			input{[]int{1, 3, 2, 3, 3}, 2},
			output{6},
		},
		{
			input{[]int{1, 4, 2, 1}, 3},
			output{0},
		},
		{
			input{[]int{1, 1, 3, 2, 3, 1, 3, 1}, 2},
			output{16},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := countSubarrays(input.nums, input.k)
		fmt.Printf(", output: %v\n", output)
	}
}
