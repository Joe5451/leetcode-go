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

func Test_Problem128(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 128 ----------")

	questions := []question{
		{
			input{[]int{100, 4, 200, 1, 3, 2}},
			output{4},
		},
		{
			input{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}},
			output{9},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := longestConsecutive(input.nums)
		fmt.Printf(", output: %v\n", output)
	}
}
