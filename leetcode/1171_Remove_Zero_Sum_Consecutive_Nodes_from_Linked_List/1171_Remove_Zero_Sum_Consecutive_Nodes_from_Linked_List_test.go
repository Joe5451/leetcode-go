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
	head []int
}

type output struct {
	result []int
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func intsToList(nums []int) *ListNode {
	headPreviousNode := &ListNode{}
	currentNode := headPreviousNode
	for _, num := range nums {
		currentNode.Next = &ListNode{Val: num}
		currentNode = currentNode.Next
	}
	return headPreviousNode.Next
}

func listToInts(list *ListNode) []int {
	nums := []int{}
	for list != nil {
		nums = append(nums, list.Val)
		list = list.Next
	}

	return nums
}

func Test_Problem1171(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 1171 ----------")

	questions := []question{
		{
			input{[]int{1, 2, -3, 3, 1}},
			output{[]int{3, 1}},
		},
		{
			input{[]int{1, 2, 3, -3, 4}},
			output{[]int{1, 2, 4}},
		},
		{
			input{[]int{1, 2, 3, -3, -2}},
			output{[]int{1}},
		},
		{
			input{[]int{1, 3, 2, -3, -2, 5, 5, -5, 1}},
			output{[]int{1, 5, 1}},
		},
		{
			input{[]int{0, 0}},
			output{[]int{}},
		},
		{
			input{[]int{2, 2, -2, 1, -1, -1}},
			output{[]int{2, -1}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := removeZeroSumSublists(intsToList(input.head))
		fmt.Printf(", output: %v\n", listToInts(output))
	}
}
