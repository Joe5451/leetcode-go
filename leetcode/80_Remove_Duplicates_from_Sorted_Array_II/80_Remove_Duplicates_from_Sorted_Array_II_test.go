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

func Test_Problem80(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 80 ----------")

	questions := []question{
		{
			input{[]int{1, 1, 1, 2, 2, 3}},
			output{5},
		},
		{
			input{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}},
			output{7},
		},
	}

	fmt.Println("solution1:")
	for _, question := range questions {
		nums := question.input.nums
		fmt.Printf("input: %v", question.input)
		output := removeDuplicates(nums)
		fmt.Printf(", output: %v\n", output)
	}

	questions = []question{
		{
			input{[]int{1, 1, 1, 2, 2, 3}},
			output{5},
		},
		{
			input{[]int{0, 0, 1, 1, 1, 1, 2, 3, 3}},
			output{7},
		},
	}

	fmt.Println("\nsolution2:")
	for _, question := range questions {
		nums := question.input.nums
		fmt.Printf("input: %v", question.input)
		output := removeDuplicates2(nums)
		fmt.Printf(", output: %v\n", output)
	}
}
