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
	obstacleGrid [][]int
}

type output struct {
	result int
}

func Test_Problem63(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 63 ----------")

	questions := []question{
		{
			input{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
			output{2},
		},
		{
			input{[][]int{{0, 1}, {0, 0}}},
			output{1},
		},
		{
			input{[][]int{{1}}},
			output{0},
		},
		{
			input{[][]int{{0}}},
			output{1},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := uniquePathsWithObstacles(input.obstacleGrid)
		fmt.Printf(", output: %v\n", output)
	}
}
