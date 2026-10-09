package twoheaps

import (
	"container/heap"
	"strconv"
)

// Structure for MaxHeap
type MaxHeap []int

// newMaxHeap function intializes an instance of the MaxHeap
func newMaxHeap() *MaxHeap {
	max := &MaxHeap{}
	heap.Init(max)
	return max
}

// Len function returns the length of the MaxHeap
func (h MaxHeap) Len() int {
	return len(h)
}

// Empty returns true if the MaxHeap is empty, false otherwise
func (h MaxHeap) Empty() bool {
	return len(h) == 0
}

// Greater returns true if the first of the given elements is greater than the second one
func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

// Swap function swaps the values at the given indices
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Top function returns the element at the top of the MaxHeap
func (h MaxHeap) Top() int {
	return h[0]
}

// Push function inserts the element in the MaxHeap
func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

// Pop function pops the element at the top of the MaxHeap
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func largestInteger(num int) int {

	// Iterate through the digits of the number and store them in a slice
	digits := []int{}
	numStr := strconv.Itoa(num)

	for _, ch := range numStr {
		digit, _ := strconv.Atoi(string(ch))
		digits = append(digits, digit)
	}

	// Create two heaps to store the odd and even digits
	oddHeap := newMaxHeap()
	evenHeap := newMaxHeap()

	for _, digit := range digits {
		if digit%2 == 0 {
			heap.Push(evenHeap, digit)
		} else {
			heap.Push(oddHeap, digit)
		}
	}

	// Iterate through the digits again and replace them with the largest available digit from the corresponding heap
	result := make([]int, len(digits))
	for i, digit := range digits {
		if digit%2 == 0 {
			result[i] = heap.Pop(evenHeap).(int)
		} else {
			result[i] = heap.Pop(oddHeap).(int)
		}
	}

	outputNumberStr := ""
	for _, d := range result {
		outputNumberStr += strconv.Itoa(d)
	}
	outputNumber, _ := strconv.Atoi(outputNumberStr)
	return outputNumber
}
