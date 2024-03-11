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
	preorder []int
	inorder []int
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

func Test_Problem105(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 105 ----------")

	questions := []question{
		{
			input{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}},
			output{[]int{3, 9, 20, NULL, NULL, 15, 7}},
		},
		{
			input{[]int{-1}, []int{-1}},
			output{[]int{-1}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := buildTree(input.preorder, input.inorder)
		fmt.Printf(", output: %v\n", Tree2ints(output))
	}
}
