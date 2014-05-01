package main 

import (
	"fmt"
	"container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int                 { return len(h) }
func (h IntHeap) Less(i, j int) bool       { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)            { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	var flowerCount int
	var customerCount int
	var price int

	fmt.Scanf("%d", &flowerCount)
	fmt.Scanf("%d", &customerCount)
	h := &IntHeap{}
	heap.Init(h)

	for i := 0; i < flowerCount; i++ {
		fmt.Scanf("%d", &price)
		heap.Push(h, price)
	}

	minimumTotal := 0
	x := 0
	for i := 0; h.Len() > 0; i++ {		
		minimumTotal += (x + 1) * heap.Pop(h).(int)
		if (i + 1) % customerCount == 0 {
			x += 1
		}
	}
	fmt.Println(minimumTotal)
}