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
	result bool
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

func Tree2ints(tn *TreeNode) []int {
	res := make([]int, 0, 984)

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

func Test_Problem98(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 98 ----------")

	questions := []question{
		{
			input{[]int{2, 1, 3}},
			output{true},
		},
		{
			input{[]int{5, 1, 4, NULL, NULL, 3, 6}},
			output{false},
		},
		{
			input{[]int{2, 2, 2}},
			output{false},
		},
		{
			input{[]int{5, 4, 6, NULL, NULL, 3, 7}},
			output{false},
		},

	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := isValidBST(Ints2TreeNode(input.root))
		fmt.Printf(", output: %v\n", output)
	}
}
