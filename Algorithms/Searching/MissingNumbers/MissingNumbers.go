// Here is O(kN) solution
// My O(nlogn) solution failed one big test case
// Tips: leverage the fact X(max) - X(min) < 100
package main

import (
    "bufio"
    "fmt"
    "os"
)

const (
	MaxValue = 10000
	MinValue = -10000
	Sentinel = -10001
)

func main() {
    var inputSize int
    io := bufio.NewReader(os.Stdin)

    fmt.Fscan(io, &inputSize)
    var listA = make([]int, inputSize)
    for i := 0; i < inputSize; i++ {
        fmt.Fscan(io, &listA[i])
    }

    min := MaxValue
    fmt.Fscan(io, &inputSize)
    var listB = make([]int, inputSize)
    for i := 0; i < inputSize; i++ {
        fmt.Fscan(io, &listB[i])
        if listB[i] < min {
        	min = listB[i]
        }
    }
    
    var orderedListB [101]int
    var countListB [101]int 
    for i := 0; i < 101; i++ {
    	orderedListB[i] = Sentinel
    }
    for _, number := range listB {
    	orderedListB[number - min] = number
    	countListB[number - min]++
    }

    for _, number := range listA {
    	countListB[number - min]--
    }

    for i := 0; i < len(countListB); i++ {
    	if countListB[i] > 0 && countListB[i] != Sentinel {
    		fmt.Print(orderedListB[i])
    		fmt.Print(" ")
    	}
    }
}


