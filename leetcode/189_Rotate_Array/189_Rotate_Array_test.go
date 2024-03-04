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
	nums []int
}

func Test_Problem189(t *testing.T) {
	questions := []question {
		{
			input{[]int{1, 2, 3, 4, 5, 6, 7}, 3},
			output{[]int{5, 6, 7, 1, 2, 3, 4}},
		},
		{
			input{[]int{-1, -100, 3, 99}, 2},
			output{[]int{3, 99, -1, -100}},
		},
	}

	fmt.Println("---------- LeetCode Problem 189 ----------")

	for _, question := range questions {
		nums := question.input.nums
		k := question.input.k
		fmt.Printf("input: %v", question.input)
		rotate(nums, k)
		fmt.Printf(", output: %v\n", nums)
	}
}
