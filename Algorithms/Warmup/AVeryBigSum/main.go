package main

import (
	"fmt"
)

func main() {
	var n int
	var sum int
	fmt.Scan(&n)
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
		sum += array[i]
	}

	fmt.Printf("%d\n", sum)
}
