package main 

import "fmt"

func main() {
	var integerCount int 
	var integer int 

	fmt.Scanf("%d", &integerCount)
	var integerArray = make([]int, 101)
	for i := 0; i < integerCount; i++ {
		fmt.Scanf("%d", &integer)
		integerArray[integer]++
	}
	for i, val := range integerArray {
		if val == 1 {
			fmt.Println(i)
			break
		}
	}
}