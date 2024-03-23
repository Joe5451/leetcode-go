package leetcode

import "container/heap"

// time complexity is O(nlogk), space complexity is O(k)
func findKthLargest(nums []int, k int) int {
	minHeap := &MinHeap{}
	for _, num := range nums {
		if minHeap.Len() < k || num > (*minHeap)[0] {
			heap.Push(minHeap, num)
			if minHeap.Len() > k {
				heap.Pop(minHeap)
			}
		}
	}

	return (*minHeap)[0]
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
