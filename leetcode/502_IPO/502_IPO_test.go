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
	k int
	w int
	profits []int
	capital []int
}

type output struct {
	result int
}

func Test_Problem502(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 502 ----------")

	questions := []question{
		{
			input{2, 0, []int{1, 2, 3}, []int{0, 1, 1}},
			output{4},
		},
		{
			input{3, 0, []int{1, 2, 3}, []int{0, 1, 2}},
			output{6},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := findMaximizedCapital(input.k, input.w, input.profits, input.capital)
		fmt.Printf(", output: %v\n", output)
	}
}
