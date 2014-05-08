package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    var numbers [101]int
    for _, number := range numbers {
        fmt.Println(number)
    }

    var inputSize int
    io := bufio.NewReader(os.Stdin)

    fmt.Fscan(io, &inputSize)
    for i := 0; i < inputSize; i++ {
        var a int
        fmt.Fscan(io, &a)

        //logic here
    }

    fmt.Fscan(io, &inputSize)
    for i := 0; i < inputSize; i++ {
        var a int
        fmt.Fscan(io, &a)

        // logic here
    }
    // logic here
}


