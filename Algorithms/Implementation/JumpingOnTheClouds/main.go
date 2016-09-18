package main

import (
	"fmt"
)

func main() {
	var n int
	var counter int
	fmt.Scan(&n)
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}
	for i := 0; i < n; i++ {
		if i+1 == n-1 || i+2 == n-1 {
			counter++
			fmt.Printf("i = %d; c = %d; n = %d WINNER\n", i, counter, n)
			break
		}
		if array[i+2] == 0 {
			counter++
			i++
			fmt.Printf("i = %d; c = %d; n = %d; +2\n", i, counter, n)
			continue
		} else if array[i+1] == 0 {
			counter++
			fmt.Printf("i = %d; c = %d; n = %d +1\n", i, counter, n)
			continue
		}
	}

	fmt.Printf("%d\n", counter)
}
