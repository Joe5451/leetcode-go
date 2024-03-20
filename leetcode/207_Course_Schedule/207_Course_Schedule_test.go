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
	numCourses int
	nums [][]int
}

type output struct {
	result bool
}

func Test_Problem207(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 207 ----------")

	questions := []question{
		{
			input{2, [][]int{{1, 0}}},
			output{true},
		},
		{
			input{2, [][]int{{1, 0}, {0, 1}}},
			output{false},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := canFinish(input.numCourses, input.nums)
		fmt.Printf(", output: %v\n", output)
	}
}
