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
	k int
}

func Test_Problem26(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 26 ----------")

	questions := []question{
		{
			input{[]int{1, 1, 2}},
			output{2},
		},
		{
			input{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}},
			output{5},
		},
	}

	for _, question := range questions {
		nums := question.input.nums
		fmt.Printf("input: %v", question.input)
		output := removeDuplicates(nums)
		fmt.Printf(", output: %v\n", output)
	}
}
