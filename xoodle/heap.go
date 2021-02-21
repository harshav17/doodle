package xoodle

import (
	"container/heap"
	"fmt"
)

// An intHeap is a min-heap of ints.
type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers
	// because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	heapSize := len(*h)
	lastNode := (*h)[heapSize-1]
	*h = (*h)[:heapSize-1]
	return lastNode
}

// Run2 practices a heap
func Run2() {
	h := &intHeap{10, 99, 7, 16, 5}
	heap.Init(h)
	heap.Push(h, "blue")
	fmt.Printf("minimum: %d\n", (*h)[0])
	// minimum: 3

	// Keep popping the minimum element
	for h.Len() > 1 {
		fmt.Printf("%d :", (*h)[len(*h)-1])
		fmt.Printf("%d :", (*h)[0])
		fmt.Printf("%d \n", heap.Pop(h))
	}
	// 3 5 7 10 16 99
}

type heapInt interface {
}
