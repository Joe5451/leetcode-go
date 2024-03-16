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
	head []int
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

func Test_Problem148(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 148 ----------")

	questions := []question{
		{
			input{[]int{4, 2, 1, 3}},
			output{[]int{1, 2, 3, 4}},
		},
		{
			input{[]int{-1, 5, 3, 4, 0}},
			output{[]int{-1, 0, 3, 4, 5}},
		},
		{
			input{[]int{}},
			output{[]int{}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := sortList(intsToList(input.head))
		fmt.Printf(", output: %v\n", listToInts(output))
	}
}
