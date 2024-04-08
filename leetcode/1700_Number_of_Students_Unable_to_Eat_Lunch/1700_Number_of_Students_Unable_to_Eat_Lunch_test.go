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
	students []int
	sandwiches []int
}

type output struct {
	result int
}

func Test_Problem1700(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 1700 ----------")

	questions := []question{
		{
			input{
				[]int{1, 1, 0, 0},
				[]int{0, 1, 0, 1},
			},
			output{0},
		},
		{
			input{
				[]int{1, 1, 1, 0, 0, 1},
				[]int{1, 0, 0, 0, 1, 1},
			},
			output{3},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := countStudents(input.students, input.sandwiches)
		fmt.Printf(", output: %v\n", output)
	}
}
