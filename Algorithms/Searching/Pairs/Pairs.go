// O(nlogn) run time
package main

import (
    "bufio"
    "fmt"
    "os"
    "container/heap"
)

type IntHeap []int

func (h IntHeap) Len() int                 { return len(h) }
func (h IntHeap) Less(i, j int) bool       { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)            { h[i], h[j] = h[j], h[i] }

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
    h := &IntHeap{}
    heap.Init(h)
    dict := map[int]bool{}
    var inputSize int
    var diff int
    io := bufio.NewReader(os.Stdin)

    fmt.Fscan(io, &inputSize)
    fmt.Fscan(io, &diff)
    for i := 0; i < inputSize; i++ {
        var number int
        fmt.Fscan(io, &number)
        heap.Push(h, number)
        dict[number] = true
    }
    pairCount := 0
    for i := 0; h.Len() > 0; i++ {
        curr := heap.Pop(h).(int)
        _, ok := dict[curr + diff]
        if ok {
            pairCount++
        }
    }
    fmt.Println(pairCount)
}


