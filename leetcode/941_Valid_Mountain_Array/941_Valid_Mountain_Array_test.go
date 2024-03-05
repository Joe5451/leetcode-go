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
	arr []int
}

type output struct {
	isMountainArray bool
}

func Test_validMountainArray(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 941 ----------")

	questions := []question{
		{
			input{[]int{2, 1}},
			output{false},
		},
		{
			input{[]int{3, 5, 5}},
			output{false},
		},
		{
			input{[]int{0, 3, 2, 1}},
			output{true},
		},
	}

	fmt.Println("solution1:")

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := validMountainArray(input.arr)
		fmt.Printf(", output: %v\n", output)
	}

	fmt.Println("\nsolution2:")

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := validMountainArray2(input.arr)
		fmt.Printf(", output: %v\n", output)
	}
}
