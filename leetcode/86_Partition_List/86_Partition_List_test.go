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
	x int
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

func Test_Problem86(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 86 ----------")

	questions := []question{
		{
			input{[]int{1, 4, 3, 2, 5, 2}, 3},
			output{[]int{1, 2, 2, 4, 3, 5}},
		},
		{
			input{[]int{2, 1}, 2},
			output{[]int{1, 2}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := partition(intsToList(input.head), input.x)
		fmt.Printf(", output: %v\n", listToInts(output))
	}
}
