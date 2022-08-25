package main

import (
	"container/heap"
	"fmt"
)

type Heap []int

// implement Len() Less() Swap() for sort.Interface
func (h Heap) Len() int {
	return len(h)
}

// max heap
func (h Heap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// implement Push() Pop() for heap interface
func (h *Heap) Push(item interface{}) {
	*h = append(*h, item.(int))
}

func (h *Heap) Pop() interface{} {
	item := (*h)[h.Len()-1]
	*h = (*h)[0 : h.Len()-1]
	return item
}

func main() {
	h := Heap{1, 3, 5, 2, 1, 0, 99}
	heap.Init(&h)

	for h.Len() > 0 {
		fmt.Println(heap.Pop(&h))
	}
}
