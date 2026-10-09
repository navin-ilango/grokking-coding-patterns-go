package topk

import (
	"container/heap"
	"sort"
)

// IntHeap defines a min-heap of integers
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // Min-heap property
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func thirdMax(nums []int) int {

	h := &IntHeap{}
	heap.Init(h)
	seen := make(map[int]bool)

	sort.Slice(nums, func(i, j int) bool { return nums[i] > nums[j] })

	for _, num := range nums {
		if !seen[num] {
			seen[num] = true
			heap.Push(h, num)
			if h.Len() > 3 {
				heap.Pop(h)
			}
		}
	}

	if h.Len() < 3 {
		return nums[0]
	}

	// The root of the min-heap of size 3 will be the 3rd maximum number
	return heap.Pop(h).(int)
}
