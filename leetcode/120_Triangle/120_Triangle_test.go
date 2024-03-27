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
	triangle [][]int
}

type output struct {
	result int
}

func Test_Problem120(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 120 ----------")

	questions := []question{
		{
			input{[][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}},
			output{11},
		},
		{
			input{[][]int{{-10}}},
			output{-10},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := minimumTotal(input.triangle)
		fmt.Printf(", output: %v\n", output)
	}
}
