package leetcode

import "container/heap"

// time complexity is O(nlogn), space complexity is O(n)
func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	capitalHeap, profitHeap := &CapitalHeap{}, &ProfitHeap{}
	result := w
	// time complexity: O(n)
	for i, cap := range capital {
		if cap <= w {
			heap.Push(profitHeap, profits[i])
		} else {
			heap.Push(capitalHeap, Item{profit: profits[i], capital: cap})
		}
	}

	// time complexity: O(k*nlogn) -> O(nlogn)
	for profitHeap.Len() > 0 && k > 0 {
		result += heap.Pop(profitHeap).(int)
		k--
		if k <= 0 {
			break
		}

		// time complexity: O(nlogn)
		for capitalHeap.Len() > 0 && result >= (*capitalHeap)[0].capital {
			heap.Push(profitHeap, heap.Pop(capitalHeap).(Item).profit)
		}
	}

	return result
}

type Item struct {
	profit int
	capital int
}

// min heap
type CapitalHeap []Item

func (h CapitalHeap) Len() int           { return len(h) }
func (h CapitalHeap) Less(i, j int) bool { return h[i].capital < h[j].capital }
func (h CapitalHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h CapitalHeap) Empty() bool        { return len(h) == 0 }

func (h *CapitalHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *CapitalHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// max heap
type ProfitHeap []int

func (h ProfitHeap) Len() int           { return len(h) }
func (h ProfitHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h ProfitHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h ProfitHeap) Empty() bool        { return len(h) == 0 }

func (h *ProfitHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *ProfitHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
