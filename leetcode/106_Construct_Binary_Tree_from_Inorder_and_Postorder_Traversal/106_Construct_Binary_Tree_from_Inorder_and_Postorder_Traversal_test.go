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
	inorder []int
	postorder []int
}

type output struct {
	result []int
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var NULL = -1 << 63

func Tree2ints(tn *TreeNode) []int {
	res := make([]int, 0, 1024)

	queue := []*TreeNode{tn}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			nd := queue[i]
			if nd == nil {
				res = append(res, NULL)
			} else {
				res = append(res, nd.Val)
				queue = append(queue, nd.Left, nd.Right)
			}
		}
		queue = queue[size:]
	}

	i := len(res)
	for i > 0 && res[i-1] == NULL {
		i--
	}

	return res[:i]
}

func Test_Problem106(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 106 ----------")

	questions := []question{
		{
			input{[]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}},
			output{[]int{3, 9, 20, NULL, NULL, 15, 7}},
		},
		{
			input{[]int{-1}, []int{-1}},
			output{[]int{-1}},
		},
		{
			input{[]int{1, 2, 3}, []int{3, 2, 1}},
			output{[]int{1, NULL, 2, NULL, 3}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := buildTree(input.inorder, input.postorder)
		fmt.Printf(", output: %v\n", Tree2ints(output))
	}
}
