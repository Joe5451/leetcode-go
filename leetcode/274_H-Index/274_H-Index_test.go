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
	citations []int
}

type output struct {
	h int
}

func Test_validMountainArray(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 274 ----------")

	questions := []question{
		{
			input{[]int{3, 0, 6, 1, 5}},
			output{3},
		},
		{
			input{[]int{1, 3, 1}},
			output{1},
		},
		{
			input{[]int{100}},
			output{1},
		},
		{
			input{[]int{0}},
			output{0},
		},
		{
			input{[]int{11, 15}},
			output{2},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := hIndex(input.citations)
		fmt.Printf(", output: %v\n", output)
	}
}
