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
	ans [][]int
}

func Test_Problem3(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 3 ----------")

	questions := []question{
		{
			input{[]int{-1, 0, 1, 2, -1, -4}},
			output{[][]int{{-1, -1, 2}, {-1, 0, 1}}},
		},
		{
			input{[]int{0, 1, 1}},
			output{[][]int{}},
		},
		{
			input{[]int{0, 0, 0}},
			output{[][]int{{0, 0, 0}}},
		},
	}

	for _, question := range questions {
		nums := question.input.nums
		fmt.Printf("input: %v", question.input)
		output := threeSum(nums)
		fmt.Printf(", output: %v\n", output)
	}
}
