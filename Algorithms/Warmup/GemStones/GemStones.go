package main 

import (
	"fmt"
	"strings"
)

func main() {
	var rockSize int
	fmt.Scanf("%d", &rockSize)
	var gemElemSet = make(map[string]bool)
	for i := 0; i < rockSize; i++ {
		var input string
		fmt.Scanf("%s", &input)
		if i == 0 {
			for _, j := range input {
				_, ok := gemElemSet[string(j)]
				if !ok {
					gemElemSet[string(j)] = true
				}
			}
		} else {
			for element, _ := range gemElemSet {
				if !strings.Contains(input, element) {
					delete(gemElemSet, element)
				}
			}
		}
	}
	fmt.Println(len(gemElemSet))

}