package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// time complexity is O(n), space complexity is O(1)
func reorderList(head *ListNode)  {
	prev, start, end := head, head, head
	for end != nil && end.Next != nil {
		tempStart := start.Next
		tempEnd := end.Next.Next
		start.Next = prev
		prev = start
		start = tempStart
		end = tempEnd
	}

	next := start
	if end != nil {
		next = start.Next
	}
	prevEnd := start
	end = start
	for next != nil {
		tempNext := next.Next
		tempPrev := prev.Next

		prev.Next = next
		next.Next = prevEnd

		prevEnd = prev
		prev = tempPrev
		next = tempNext
	}

	end.Next = nil
}
