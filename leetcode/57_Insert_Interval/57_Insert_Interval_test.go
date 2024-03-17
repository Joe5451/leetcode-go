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
	intervals [][]int
	newInterval []int
}

type output struct {
	result [][]int
}

func Test_Problem57(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 57 ----------")

	questions := []question{
		{
			input{[][]int{{1, 3}, {6, 9}}, []int{2, 5}},
			output{[][]int{{1, 5}, {6, 9}}},
		},
		{
			input{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}},
			output{[][]int{{1, 2}, {3, 10}, {12, 16}}},
		},
		{
			input{[][]int{{1, 3}, {6, 9}}, []int{4, 5}},
			output{[][]int{{1, 3}, {4, 5}, {6, 9}}},
		},
		{
			input{[][]int{}, []int{5, 7}},
			output{[][]int{{5, 7}}},
		},
		{
			input{[][]int{{1, 5}}, []int{2, 7}},
			output{[][]int{{1, 7}}},
		},
		{
			input{[][]int{{1, 5}}, []int{6, 8}},
			output{[][]int{{1, 5}, {6, 8}}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := insert(input.intervals, input.newInterval)
		fmt.Printf(", output: %v\n", output)
	}
}
