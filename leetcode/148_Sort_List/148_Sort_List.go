package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// time complexity is O(n logn), space complexity is O(1)
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	prev, slow, fast := head, head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}

	prev.Next = nil

	leftHead := sortList(head)
	rightHead := sortList(slow)

	return mergeList(leftHead, rightHead)
}

func mergeList(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	}

	if head2 == nil {
		return head1
	}
	dummy := &ListNode{0, nil}
	prev := dummy
	current1 := head1
	current2 := head2
	for current1 != nil && current2 != nil {
		if current1.Val < current2.Val {
			prev.Next = current1
			current1 = current1.Next
		} else {
			prev.Next = current2
			current2 = current2.Next
		}
		prev = prev.Next
	}

	if current1 != nil {
		prev.Next = current1
	}

	if current2 != nil {
		prev.Next = current2
	}

	return dummy.Next
}
