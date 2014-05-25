// O(n) solution or 0(kn) 
// with k constant to be exact.
package main 

import (
	"fmt"
	"bufio"
	"os"
)
func main() {
	var inputSize int 
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &inputSize)
	for i := 0; i < inputSize; i++ {
		var input string
		fmt.Fscan(io, &input)
		if len(input) % 2 != 0 {
			fmt.Println(-1)
		} else {
			// put inputs into counting map
			// note the conversion from rune to ascii
			// because go does not have char type
			// and string is not necessary as key
			firstStrMap := make(map[int]int)
			secondStrMap := make(map[int]int)
			mid := len(input) / 2
			for i, char := range input {
				utf := int(char)
				if i < mid {
					_, ok := firstStrMap[utf]
					if ok {
						firstStrMap[utf]++
					} else {
						firstStrMap[utf] = 1
					}
				} else {
					_, ok := secondStrMap[utf]
					if ok {
						secondStrMap[utf]++
					} else {
						secondStrMap[utf] = 1
					}
				}
			}

			// truncate any overlap ocurrences in two maps
			for letter, count := range secondStrMap {
				letterCountInFirstMap, ok := firstStrMap[letter]
				if ok {
					if letterCountInFirstMap == count {
						delete(firstStrMap, letter)
						delete(secondStrMap, letter)
					}
					if letterCountInFirstMap > count {
						delete(secondStrMap, letter)
						firstStrMap[letter] -= count
					}
					if letterCountInFirstMap < count {
						delete(firstStrMap, letter)
						secondStrMap[letter] -= letterCountInFirstMap
					}
				}
			}

			// count the letters needed to replace to obtain an anagram
			total := 0
			for _, count := range firstStrMap {
				total += count
			}
			fmt.Println(total)
		}
	}

}