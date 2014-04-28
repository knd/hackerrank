package main

import (
		"fmt"
		"bytes"
		"strconv"
)

func main() {
		var size int
		fmt.Scanf("%d", &size)

		var array = make([]int, size)

		for i, _ := range array {
				fmt.Scanf("%d", &array[i])
		}
		for i, element := range array {
				j := i
				for ; j > 0 && element < array[j-1]; j-- {
						array[j] = array[j-1]
				}
				array[j] = element
		}
		var strBuffer bytes.Buffer
		for _, element := range array {
				strBuffer.WriteString(strconv.Itoa(element))
				strBuffer.WriteString(" ")
		}
		fmt.Println(strBuffer.String())
}