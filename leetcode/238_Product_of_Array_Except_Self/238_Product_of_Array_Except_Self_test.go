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
	answer []int
}

func Test_Problem238(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 238 ----------")

	questions := []question{
		{
			input{[]int{1, 2, 3, 4}},
			output{[]int{24, 12, 8, 6}},
		},
		{
			input{[]int{-1, 1, 0, -3, 3}},
			output{[]int{0, 0, 9, 0, 0}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := productExceptSelf(input.nums)
		fmt.Printf(", output: %v\n", output)
	}
}
