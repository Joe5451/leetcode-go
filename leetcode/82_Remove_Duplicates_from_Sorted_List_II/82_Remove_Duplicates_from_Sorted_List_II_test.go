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
	output []int
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

func Test_Problem82(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 82 ----------")

	questions := []question{
		{
			input{[]int{1, 2, 3, 3, 4, 4, 5}},
			output{[]int{1, 2, 5}},
		},
		{
			input{[]int{1, 1, 1, 2, 3}},
			output{[]int{2, 3}},
		},
		{
			input{[]int{1, 1, 2, 2}},
			output{[]int{}},
		},
		{
			input{[]int{2, 2}},
			output{[]int{}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := deleteDuplicates(intsToList(input.head))
		fmt.Printf(", output: %v\n", listToInts(output))
	}
}
