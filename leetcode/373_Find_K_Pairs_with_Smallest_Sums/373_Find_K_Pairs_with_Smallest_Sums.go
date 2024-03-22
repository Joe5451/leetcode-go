package leetcode

import "container/heap"

// time complexity is O(n + k), space complexity is O(n)
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	result := [][]int{}
	minHeap := &MinHeap{}

	for i, num1 := range nums1 {
		heap.Push(minHeap, Pair{num1 + nums2[0], i, 0})
	}

	for minHeap.Len() > 0 && k > 0 {
		currentPair := heap.Pop(minHeap).(Pair)
		result = append(result, []int{nums1[currentPair.i], nums2[currentPair.j]})

		nextNums2Index := currentPair.j + 1
		if nextNums2Index < len(nums2) {
			sum := nums1[currentPair.i] + nums2[nextNums2Index]
			heap.Push(minHeap, Pair{sum, currentPair.i, nextNums2Index})
		}
		k--
	}

	return result
}

type MinHeap []Pair
type Pair struct {
	sum int
	i int
	j int
}

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].sum < h[j].sum
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Pair))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
