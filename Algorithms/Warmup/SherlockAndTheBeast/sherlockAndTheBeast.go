// Things learned and tips:
// have to follow coding style if-else of go lang, otherwise compile error
// bytes.Buffer is absolutely faster than str += concatenation
package main

import (
		"fmt"
		"bytes"
)	

func main() {
		var testCount int
		var digitCount int
		
		fmt.Scanf("%d", &testCount)

		for i := 0; i < testCount; i++ {
				fmt.Scanf("%d", &digitCount)
				
				found := false
				digit5Count := (digitCount / 3) * 3
				digit3Count := digitCount - digit5Count
				for ; digit3Count <= digitCount && !found; {
						if digit3Count % 5 == 0 {
								found = true
								break
						}						
						digit5Count -= 3
						digit3Count += 3						
				}
				if found {
						var strBuffer bytes.Buffer
						for j := 0; j < digit5Count; j++ {
								strBuffer.WriteString("5")
						}
						for j := 0; j < digit3Count; j++ {
								strBuffer.WriteString("3")
						}
						fmt.Println(strBuffer.String())
				} else {
						fmt.Println(-1)
				}
		}
}		