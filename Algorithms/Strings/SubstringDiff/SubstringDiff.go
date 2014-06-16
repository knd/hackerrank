// This solution is probably the best as the run time is bound to O(N^2). 
// However, it is not my solution and the credit is given to this guy right here: 
// http://www.quora.com/What-is-the-algorithmic-approach-to-solve-hackerrank-problem-Substring-Diff
// FYI, I original could get down to (N^2logN).
package main 

import (
	"fmt"
	"os"
	"bufio"
)

func getBest(firstStr, secondStr string, i, j, k int) int {
	best := 0
	bad := 0
	ip := i
	jp := j 
	l := 0
	for {
		if ip + l >= len(firstStr) || jp + l >= len(firstStr) {
			if l > best {
				best = l
			}
			break
		}

		if firstStr[ip + l] != secondStr[jp + l] {
			bad += 1
		}

		if bad > k {
			if l > best {
				best = l
			}
			for firstStr[ip] == secondStr[jp] {
				ip += 1
				jp += 1
				l -= 1
			}
			bad -= 1
			ip += 1
			jp += 1
		} else {
			l += 1
		}
	}

	return best
}

func main() {
	var inputSize int 
	io := bufio.NewReader(os.Stdin)
	fmt.Fscan(io, &inputSize)
	var maxDiff int 
	var firstStr string 
	var secondStr string 

	for i := 0; i < inputSize; i++ {
		fmt.Fscan(io, &maxDiff)
		fmt.Fscan(io, &firstStr)
		fmt.Fscan(io, &secondStr)
		best := 0
		for j := 0; j < len(firstStr); j++ {
			best1 := getBest(firstStr, secondStr, 0, j, maxDiff)
			best2 := getBest(firstStr, secondStr, j, 0, maxDiff)
			if (best1 > best && best1 >= best2) {
				best = best1
			}
			if (best2 > best && best2 >= best1) {
				best = best2
			}
		}
		fmt.Println(best)
	}

}