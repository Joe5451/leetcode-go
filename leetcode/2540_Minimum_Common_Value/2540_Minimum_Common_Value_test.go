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
	common int
}

func Test_Problem2540(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 2540 ----------")

	questions := []question{
		{
			input{[]int{1, 2, 3}, []int{2, 4}},
			output{2},
		},
		{
			input{[]int{1, 2, 3, 6}, []int{2, 3, 4, 5}},
			output{2},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := getCommon(input.nums1, input.nums2)
		fmt.Printf(", output: %v\n", output)
	}
}
