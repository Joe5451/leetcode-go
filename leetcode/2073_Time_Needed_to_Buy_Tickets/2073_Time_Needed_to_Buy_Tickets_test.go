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
	tickets []int
	k int
}

type output struct {
	result int
}

func Test_Problem2073(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 2073 ----------")

	questions := []question{
		{
			input{
				[]int{2, 3, 2},
				2,
			},
			output{6},
		},
		{
			input{
				[]int{5, 1, 1, 1},
				0,
			},
			output{8},
		},
		{
			input{
				[]int{
					15, 66, 3, 47, 71, 27, 54, 43, 97, 34, 94, 33, 54, 26, 15,
					52, 20, 71, 88, 42, 50, 6, 66, 88, 36, 99, 27, 82, 7, 72,
				},
				18,
			},
			output{1457},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := timeRequiredToBuy(input.tickets, input.k)
		fmt.Printf(", output: %v\n", output)
	}
}
