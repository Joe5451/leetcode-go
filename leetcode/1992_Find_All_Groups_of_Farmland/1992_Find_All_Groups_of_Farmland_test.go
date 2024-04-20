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
	land [][]int
}

type output struct {
	result [][]int
}

func Test_Problem1992(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 1992 ----------")

	questions := []question{
		{
			input{
				[][]int{
					{1, 0, 0},
					{0, 1, 1},
					{0, 1, 1},
				},
			},
			output{
				[][]int{
					{0, 0, 0, 0},
					{1, 1, 2, 2},
				},
			},
		},
		{
			input{
				[][]int{
					{1, 1},
					{1, 1},
				},
			},
			output{
				[][]int{
					{0, 0, 1, 1},
				},
			},
		},
		{
			input{
				[][]int{
					{0},
				},
			},
			output{
				[][]int{},
			},
		},
		{
			input{
				[][]int{
					{1, 1, 0, 0, 0, 1},
					{1, 1, 0, 0, 0, 0},
				},
			},
			output{
				[][]int{
					{0, 0, 1, 1},
					{0, 5, 0, 5},
				},
			},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := findFarmland(input.land)
		fmt.Printf(", output: %v\n", output)
	}
}
