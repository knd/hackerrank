// easy problem that keeps blood brain flowing
package main

import "fmt"	

func main() {
		var unitCount int
		var testCount int
		var entryIndex int
		var exitIndex int
		
		fmt.Scanf("%d", &unitCount)
		fmt.Scanf("%d", &testCount)

		var units = make([]int, unitCount)
		for i, _ := range units {
				fmt.Scanf("%d", &units[i])
		}

		for i := 0; i < testCount; i++ {
				fmt.Scanf("%d", &entryIndex)
				fmt.Scanf("%d", &exitIndex)
				minWidth := units[entryIndex]
				for j := entryIndex + 1; j <= exitIndex; j++ {
						if units[j] < minWidth {
								minWidth = units[j]
						}
				}
				fmt.Println(minWidth)	
		}
}		