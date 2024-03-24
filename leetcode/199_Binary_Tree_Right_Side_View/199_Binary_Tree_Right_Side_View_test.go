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
	root []int
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

func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: ints[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && ints[i] != NULL {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && ints[i] != NULL {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func Test_Problem199(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 199 ----------")

	questions := []question{
		{
			input{[]int{1, 2, 3, NULL, 5, NULL, 4}},
			output{[]int{1, 3, 4}},
		},
		{
			input{[]int{1, NULL, 3}},
			output{[]int{1, 3}},
		},
		{
			input{[]int{}},
			output{[]int{}},
		},
		{
			input{[]int{1, 2, 3, NULL, 5, 6, NULL}},
			output{[]int{1, 3, 6}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := rightSideView(Ints2TreeNode(input.root))
		fmt.Printf(", output: %v\n", output)
	}
}
