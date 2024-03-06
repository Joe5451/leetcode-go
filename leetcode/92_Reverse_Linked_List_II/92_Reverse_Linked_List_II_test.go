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
	left int
	right int
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

func Test_Problem92(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 92 ----------")

	questions := []question{
		{
			input{[]int{1, 2, 3, 4, 5}, 2, 4},
			output{[]int{1, 4, 3, 2, 5}},
		},
		{
			input{[]int{5}, 1, 1},
			output{[]int{5}},
		},
	}

	fmt.Println("solution1:")
	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := reverseBetween(intsToList(input.head), input.left, input.right)
		fmt.Printf(", output: %v\n", listToInts(output))
	}

	fmt.Println("\nsolution2:")
	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := reverseBetween2(intsToList(input.head), input.left, input.right)
		fmt.Printf(", output: %v\n", listToInts(output))
	}
}
