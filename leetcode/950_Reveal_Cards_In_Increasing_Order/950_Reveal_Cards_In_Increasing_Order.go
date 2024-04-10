package leetcode

import "sort"

// time complexity is O(n), space complexity is O(n)
func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)
	result := make([]int, len(deck))
	q := &Queue{}
	for i := 0; i < len(deck); i++ {
		q.Enqueue(i)
	}

	i := 0
	for !q.IsEmpty() {
		result[q.Dequeue()] = deck[i]
		i++

		if !q.IsEmpty() {
			q.Enqueue(q.Dequeue())
		}
	}

	return result
}

type Queue []int

func (q *Queue) Enqueue(d int) {
	*q = append(*q, d)
}

func (q *Queue) Dequeue() int {
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
