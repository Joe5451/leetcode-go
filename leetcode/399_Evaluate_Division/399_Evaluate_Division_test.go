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
	equations [][]string
	values []float64
	queries [][]string
}

type output struct {
	result []float64
}

func Test_Problem399(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 399 ----------")

	questions := []question{
		{
			input{
				[][]string{{"a", "b"}, {"b", "c"}},
				[]float64{2.0, 3.0},
				[][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}},
			},
			output{[]float64{6.0, 0.5, -1.0, 1.0, -1.0}},
		},
		{
			input{
				[][]string{{"a", "b"}, {"b", "c"}, {"bc", "cd"}},
				[]float64{1.5, 2.5, 5.0},
				[][]string{{"a", "c"}, {"c", "b"}, {"bc", "cd"}, {"cd", "bc"}},
			},
			output{[]float64{3.75000, 0.40000, 5.00000, 0.20000}},
		},
		{
			input{
				[][]string{{"a", "b"}},
				[]float64{0.5},
				[][]string{{"a", "b"}, {"b", "a"}, {"a", "c"}, {"x", "y"}},
			},
			output{[]float64{0.50000, 2.00000, -1.00000, -1.00000}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := calcEquation(input.equations, input.values, input.queries)
		fmt.Printf(", output: %v\n", output)
	}
}
