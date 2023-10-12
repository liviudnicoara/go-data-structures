package main

import (
	"container/heap"
	"fmt"
)

// Problem Statement:
// Given an array of integers, find the kth smallest element.

type MinHeap []int

// Less implements heap.Interface.
func (h MinHeap) Less(i int, j int) bool {
	return h[i] < h[j]
}

// Pop implements heap.Interface.
func (h *MinHeap) Pop() any {
	l := len(*h)
	v := (*h)[l-1]
	*h = (*h)[:l-1]
	return v
}

// Push implements heap.Interface.
func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

// Swap implements heap.Interface.
func (h MinHeap) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h MinHeap) Len() int {
	return len(h)
}

func findSmallestElem(num []int, k int) int {
	h := &MinHeap{}
	heap.Init(h)

	for _, n := range num {
		heap.Push(h, n)
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	return heap.Pop(h).(int)

}

func main() {
	// 4 is the 4 smallest element
	list := []int{9, 4, 5, 2, 3, 1, 8, 1}

	fmt.Println(findSmallestElem(list, 4))
}
