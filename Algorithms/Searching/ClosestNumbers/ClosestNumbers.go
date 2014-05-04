package main 

import (
	"fmt"
	"math"
	"container/heap"
)

type IntHeap []int 

func (h IntHeap) Len() int              { return len(h) }
func (h IntHeap) Less(i, j int) bool    { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[0 : n-1]
	return x
}

func main() {
	var numberCount int
	var number int
	h := &IntHeap{}
	heap.Init(h)

	fmt.Scanf("%d", &numberCount)
	for i := 0; i < numberCount; i++ {
		fmt.Scanf("%d", &number)
		heap.Push(h, number)
	}

	var outputNumbers []int
	prevNumber := heap.Pop(h).(int)
	currNumber := heap.Pop(h).(int)
	diff := int(math.Abs(float64(currNumber - prevNumber)))
	outputNumbers = append(outputNumbers, prevNumber, currNumber)
	prevNumber = currNumber

	for i := 0; h.Len() > 0; i++ {
		currNumber = heap.Pop(h).(int)
		newDiff := int(math.Abs(float64(currNumber - prevNumber)))
		if diff > newDiff {
			diff = newDiff
			outputNumbers = nil
			outputNumbers = append(outputNumbers, prevNumber, currNumber)
		} else if diff == newDiff {
			outputNumbers = append(outputNumbers, prevNumber, currNumber)
		}
		prevNumber = currNumber
	}

	for _, number := range outputNumbers {
		fmt.Print(number)
		fmt.Print(" ")
	}

}