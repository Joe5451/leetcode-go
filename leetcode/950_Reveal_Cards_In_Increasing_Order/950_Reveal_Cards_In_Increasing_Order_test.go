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
	deck []int
}

type output struct {
	result []int
}

func Test_Problem950(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 950 ----------")

	questions := []question{
		{
			input{[]int{17, 13, 11, 2, 3, 5, 7}},
			output{[]int{2, 13, 3, 11, 5, 17, 7}},
		},
		{
			input{[]int{1, 1000}},
			output{[]int{1, 1000}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := deckRevealedIncreasing(input.deck)
		fmt.Printf(", output: %v\n", output)
	}
}
