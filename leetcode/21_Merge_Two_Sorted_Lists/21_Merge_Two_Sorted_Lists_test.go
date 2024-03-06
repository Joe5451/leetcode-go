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
	mergedList []int
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

func Test_Problem21(t *testing.T) {
	fmt.Println("---------- LeetCode Problem 21 ----------")

	questions := []question{
		{
			input{[]int{1, 2, 4}, []int{1, 3, 4}},
			output{[]int{1, 1, 2, 3, 4, 4}},
		},
		{
			input{[]int{}, []int{}},
			output{[]int{}},
		},
		{
			input{[]int{}, []int{0}},
			output{[]int{0}},
		},
	}

	for _, question := range questions {
		input := question.input
		fmt.Printf("input: %v", input)
		mergedList := mergeTwoLists(intsToList(input.nums1), intsToList(input.nums2))
		fmt.Printf(", output: %v\n", listToInts(mergedList))
	}
}
