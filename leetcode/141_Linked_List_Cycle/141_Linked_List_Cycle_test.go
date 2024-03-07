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
	isCycle bool
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

func Test_Problem141(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 141 ----------")

	questions := []question{
		{
			input{[]int{3, 2, 0, -4}},
			output{false},
		},
		{
			input{[]int{1, 2}},
			output{false},
		},
		{
			input{[]int{1}},
			output{false},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := hasCycle(intsToList(input.head))
		fmt.Printf(", output: %v\n", output)
	}
}
