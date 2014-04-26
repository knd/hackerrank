// easy problem that keeps blood brain flowing
package main

import "fmt"	

func main() {
		var testCase int
		var cuts int
		
		fmt.Scanf("%d", &testCase)
		for i := 0; i < testCase; i++ {
				fmt.Scanf("%d", &cuts)
				mid := cuts / 2
				fmt.Println(mid * (cuts - mid))
		}
}		