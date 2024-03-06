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
	nums1 []int
	nums2 []int
}

type output struct {
	nums []int
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

func Test_Problem2(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 2 ----------")

	questions := []question{
		{
			input{[]int{2, 4, 3}, []int{5, 6, 4}},
			output{[]int{7, 0, 8}},
		},
		{
			input{[]int{0}, []int{0}},
			output{[]int{0}},
		},
		{
			input{[]int{9, 9, 9, 9, 9, 9, 9}, []int{9, 9, 9, 9}},
			output{[]int{8, 9, 9, 9, 0, 0, 0, 1}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		mergedList := addTwoNumbers(intsToList(input.nums1), intsToList(input.nums2))
		fmt.Printf(", output: %v\n", listToInts(mergedList))
	}
}
