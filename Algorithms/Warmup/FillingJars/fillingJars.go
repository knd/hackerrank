// huge test cases that the testing server
// is likely to deviate about 1 sec and solution 
// could go beyond 6 sec and fail. 
package main

import "fmt"	

func main() {
		var jars int
		var operations int
		var startIndex int
		var endIndex int
		var addedCandies int
		var totalCandies int
		
		fmt.Scanf("%d", &jars)
		fmt.Scanf("%d", &operations)
		for i := 0; i < operations; i++ {
				fmt.Scanf("%d", &startIndex)
				fmt.Scanf("%d", &endIndex)
				fmt.Scanf("%d", &addedCandies)
				totalCandies += (endIndex - startIndex + 1) * addedCandies	
		}
		fmt.Println(totalCandies / jars)
}		