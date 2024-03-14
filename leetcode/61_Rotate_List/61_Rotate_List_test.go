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
	k int
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

func Test_Problem61(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 61 ----------")

	questions := []question{
		{
			input{[]int{1, 2, 3, 4, 5}, 2},
			output{[]int{4, 5, 1, 2, 3}},
		},
		{
			input{[]int{0, 1, 2}, 4},
			output{[]int{2, 0, 1}},
		},
		{
			input{[]int{}, 0},
			output{[]int{}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := rotateRight(intsToList(input.head), input.k)
		fmt.Printf(", output: %v\n", listToInts(output))
	}
}
