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
	val int
}

type output struct {
	result int
}

func Test_Problem27(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 27 ----------")

	questions := []question{
		{
			input{[]int{3, 2, 2, 3}, 3},
			output{2},
		},
		{
			input{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2},
			output{5},
		},
		{
			input{[]int{1}, 1},
			output{0},
		},
		{
			input{[]int{4, 5}, 5},
			output{1},
		},
		{
			input{[]int{}, 0},
			output{0},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := removeElement(input.nums, input.val)
		fmt.Printf(", output: %v\n", output)
	}
}
