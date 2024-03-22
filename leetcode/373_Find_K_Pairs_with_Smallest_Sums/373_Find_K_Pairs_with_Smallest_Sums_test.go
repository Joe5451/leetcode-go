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
	k int
}

type output struct {
	result [][]int
}

func Test_Problem373(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 373 ----------")

	questions := []question{
		{
			input{[]int{1, 7, 11}, []int{2, 4, 6}, 3},
			output{[][]int{{1, 2}, {1, 4}, {1, 6}}},
		},
		{
			input{[]int{1, 1, 2}, []int{1, 2, 3}, 2},
			output{[][]int{{1, 1}, {1, 1}}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := kSmallestPairs(input.nums1, input.nums2, input.k)
		fmt.Printf(", output: %v\n", output)
	}
}
