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

func Test_Problem122(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 122 ----------")

	questions := []question{
		{
			input{[]int{7, 1, 5, 3, 6, 4}},
			output{7},
		},
		{
			input{[]int{1, 2, 3, 4, 5}},
			output{4},
		},
		{
			input{[]int{7, 6, 4, 3, 1}},
			output{0},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := maxProfit(input.prices)
		fmt.Printf(", output: %v\n", output)
	}
}
