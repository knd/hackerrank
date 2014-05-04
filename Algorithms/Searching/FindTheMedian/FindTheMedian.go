package main 

import "fmt"

func main() {
	var numberCount int

	fmt.Scanf("%d", &numberCount)
	var numberArray = make([]int, numberCount)
	for i := 0; i < numberCount; i++ {
		fmt.Scanf("%d", &numberArray[i])
	}
	median := numberCount / 2 + 1
	fmt.Println(findMedian(numberArray, median))
}


func findMedian(numberArray []int, median int) int {
	pivotNumber := numberArray[0]
	numberArray = numberArray[1:]
	var left []int
	var right []int
	for _, number := range numberArray {
		if number <= pivotNumber {
			left = append(left, number)
		} else {
			right = append(right, number)
		} 
	}
	if median == len(left) + 1 {
		return pivotNumber
	} else if median < len(left) + 1 {
		return findMedian(left, median)
	} else {
		return findMedian(right, median - len(left) - 1)
	}
	panic("To surpress compile error in go lang.")
}