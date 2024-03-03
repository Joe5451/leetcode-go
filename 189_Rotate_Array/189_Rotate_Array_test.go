package leetcode

import (
	"testing"
	"fmt"
)

type question189 struct {
	input189
	output189
}

type input189 struct {
	nums []int
	k int
}

type output189 struct {
	nums []int
}

func Test_Problem189(t *testing.T) {
	questions := []question189 {
		{
			input189{[]int{1, 2, 3, 4, 5, 6, 7}, 3},
			output189{[]int{5, 6, 7, 1, 2, 3, 4}},
		},
		{
			input189{[]int{-1, -100, 3, 99}, 2},
			output189{[]int{3, 99, -1, -100}},
		},
	}

	fmt.Println("---------- LeetCode Problem 189 ----------")

	for _, question := range questions {
		nums := question.input189.nums
		k := question.input189.k
		fmt.Printf("input: %v", question.input189)
		rotate(nums, k)
		fmt.Printf(", output: %v\n", nums)
	}
}
