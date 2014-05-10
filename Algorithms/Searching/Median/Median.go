// worst case O(nlogn) for add operation,
// worst case O((n/2)^2) for remove operation,
// because number is 32-bit signed, it's possbily 
// that we encounter -2^31 to 2^31-1, then the 
// summation of 2 numbers could take more than 4 bytes
// and so it's safe best to hold inputs in int64.
package main

import (
    "bufio"
    "fmt"
    "os"
    "container/heap"
    "errors"
    "math"
)

/*-----------------------------------------------------------------*/
type Int64MinHeap []int64
func (h Int64MinHeap)  Len() int                { return len(h) }
func (h Int64MinHeap)  Less(i, j int) bool      { return h[i] < h[j] }
func (h Int64MinHeap)  Swap(i, j int)           { h[i], h[j] = h[j], h[i] }
func (h *Int64MinHeap) Push(x interface{})      { *h = append(*h, x.(int64)) }
func (h *Int64MinHeap) Pop() interface{} {
    n := len(*h)
    x := (*h)[n-1]
    *h = (*h)[0 : n-1]
    return x
}

type Int64MaxHeap []int64
func (h Int64MaxHeap)  Len() int                { return len(h) }
func (h Int64MaxHeap)  Less(i, j int) bool      { return h[i] > h[j] }
func (h Int64MaxHeap)  Swap(i, j int)           { h[i], h[j] = h[j], h[i] }
func (h *Int64MaxHeap) Push(x interface{})      { *h = append(*h, x.(int64)) }
func (h *Int64MaxHeap) Pop() interface{} {
    n := len(*h)
    x := (*h)[n-1]
    *h = (*h)[0 : n-1]
    return x
}
/*-----------------------------------------------------------------*/

func removeFromMinHeap(mih *Int64MinHeap, value int64) error {
    for i := 0; i < len(*mih); i++ {
        if (*mih)[i] == value {
            heap.Remove(mih, i)
            return nil
        }
    }
    return errors.New("Element does not exist or empty heap.")
}

func removeFromMaxHeap(mah *Int64MaxHeap, value int64) error {
    for i := 0; i < len(*mah); i++ {
        if (*mah)[i] == value {
            heap.Remove(mah, i)
            return nil
        }
    }
    return errors.New("Element does not exist or empty heap.")
}


func main() {
    lowerHalfMaxHeap := &Int64MaxHeap{}
    upperHalfMinHeap := &Int64MinHeap{}
    heap.Init(lowerHalfMaxHeap)
    heap.Init(upperHalfMinHeap)

    var inputSize int
    io := bufio.NewReader(os.Stdin)
    fmt.Fscan(io, &inputSize)
    for i := 0; i < inputSize; i++ {
        var op string
        var value int64
        var removeErr error
        fmt.Fscan(io, &op)
        fmt.Fscan(io, &value)

        // add or remove element from either heap
        switch op {
        case "a":
            if lowerHalfMaxHeap.Len() == 0 {
                heap.Push(lowerHalfMaxHeap, value)
            } else if (*lowerHalfMaxHeap)[0] >= value {
                heap.Push(lowerHalfMaxHeap, value)
            } else {
                heap.Push(upperHalfMinHeap, value)
            }
        case "r":
            if lowerHalfMaxHeap.Len() == 0 {
                removeErr = errors.New("No element to remove.")
            } else {
                if (*lowerHalfMaxHeap)[0] >= value {
                    removeErr = removeFromMaxHeap(lowerHalfMaxHeap, value)
                } else {
                    removeErr = removeFromMinHeap(upperHalfMinHeap, value)
                }
            }

        }

        // balance the length of both heaps that the difference
        // between them never exceeds one element
        if upperHalfMinHeap.Len() > lowerHalfMaxHeap.Len() {
            val := heap.Pop(upperHalfMinHeap).(int64)
            heap.Push(lowerHalfMaxHeap, val)
        }
        if upperHalfMinHeap.Len() + 1 < lowerHalfMaxHeap.Len() {
            val := heap.Pop(lowerHalfMaxHeap).(int64)
            heap.Push(upperHalfMinHeap, val)
        }

        // calculate the median and output to console
        if removeErr != nil || lowerHalfMaxHeap.Len() == 0 {
            fmt.Println("Wrong!")
        } else if lowerHalfMaxHeap.Len() == upperHalfMinHeap.Len() {
            sum := (*lowerHalfMaxHeap)[0] + (*upperHalfMinHeap)[0]
            median := sum / 2
            if int64(math.Abs(float64(sum))) % 2 == 0 {
                fmt.Println(median)
            } else {
                if sum == -1 {
                    fmt.Println("-0.5")
                } else {
                    fmt.Printf("%d.5\n", median)
                }
            }
        } else {
            fmt.Println((*lowerHalfMaxHeap)[0])
        }
    }
}


