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
	startGene string
	endGene string
	bank []string
}

type output struct {
	result int
}

func Test_Problem433(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 433 ----------")

	questions := []question{
		{
			input{"AACCGGTT", "AACCGGTA", []string{"AACCGGTA"}},
			output{1},
		},
		{
			input{"AACCGGTT", "AAACGGTA", []string{"AACCGGTA", "AACCGCTA", "AAACGGTA"}},
			output{2},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := minMutation(input.startGene, input.endGene, input.bank)
		fmt.Printf(", output: %v\n", output)
	}
}
