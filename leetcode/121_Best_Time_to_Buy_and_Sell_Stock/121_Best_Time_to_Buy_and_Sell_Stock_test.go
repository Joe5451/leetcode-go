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
	prices []int
}

type output struct {
	profit int
}

func Test_Problem121(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 121 ----------")

	questions := []question{
		{
			input{[]int{7, 1, 5, 3, 6, 4}},
			output{5},
		},
		{
			input{[]int{7, 6, 4, 3, 1}},
			output{0},
		},
		{
			input{[]int{2, 1, 2, 1, 0, 1, 2}},
			output{2},
		},
		{
			input{[]int{3, 3, 5, 0, 0, 3, 1, 4}},
			output{4},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := maxProfit(input.prices)
		fmt.Printf(", output: %v\n", output)
	}
}
