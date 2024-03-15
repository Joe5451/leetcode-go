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
	gas []int
	cost []int
}

type output struct {
	start int
}

func Test_Problem134(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 134 ----------")

	questions := []question{
		{
			input{[]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}},
			output{3},
		},
		{
			input{[]int{2, 3, 4}, []int{3, 4, 3}},
			output{-1},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := canCompleteCircuit(input.gas, input.cost)
		fmt.Printf(", output: %v\n", output)
	}
}
