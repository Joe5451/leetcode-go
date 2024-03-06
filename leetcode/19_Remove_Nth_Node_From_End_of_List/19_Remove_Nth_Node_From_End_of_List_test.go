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
	n int
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

func Test_Problem19(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 19 ----------")

	questions := []question{
		{
			input{[]int{1, 2, 3, 4, 5}, 2},
			output{[]int{1, 2, 3, 5}},
		},
		{
			input{[]int{1}, 1},
			output{[]int{}},
		},
		{
			input{[]int{1, 2}, 1},
			output{[]int{1}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := removeNthFromEnd(intsToList(input.head), input.n)
		fmt.Printf(", output: %v\n", listToInts(output))
	}
}
