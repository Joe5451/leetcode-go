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
	list1 []int
	a int
	b int
	list2 []int
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

func Test_Problem1669(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 1669 ----------")

	questions := []question{
		{
			input{[]int{10, 1, 13, 6, 9, 5}, 3, 4, []int{1000000, 1000001, 1000002}},
			output{[]int{10, 1, 13, 1000000, 1000001, 1000002, 5}},
		},
		{
			input{[]int{0, 1, 2, 3, 4, 5, 6}, 2, 5, []int{1000000, 1000001, 1000002, 1000003, 1000004}},
			output{[]int{0, 1, 1000000, 1000001, 1000002, 1000003, 1000004, 6}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		output := mergeInBetween(intsToList(input.list1), input.a, input.b, intsToList(input.list2))
		fmt.Printf(", output: %v\n", listToInts(output))
	}
}
