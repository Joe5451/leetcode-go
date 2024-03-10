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
	nums2 []int
}

type output struct {
	result []int
}

func Test_Problem349(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 349 ----------")

	questions := []question{
		{
			input{[]int{1, 2, 2, 1}, []int{2, 2}},
			output{[]int{2}},
		},
		{
			input{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}},
			output{[]int{9, 4}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := intersection(input.nums1, input.nums2)
		fmt.Printf(", output: %v\n", output)
	}
}
