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
	nums1 []int
	m int
	nums2 []int
	n int
}

type output struct {
	nums1 []int
}

func Test_Problem88(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 88 ----------")

	questions := []question{
		{
			input{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
			output{[]int{1, 2, 2, 3, 5, 6}},
		},
		{
			input{[]int{1}, 1, []int{}, 0},
			output{[]int{1}},
		},
		{
			input{[]int{0}, 0, []int{1}, 1},
			output{[]int{1}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		merge(input.nums1, input.m, input.nums2, input.n)
		fmt.Printf(", output: %v\n", input.nums1)
	}
}
