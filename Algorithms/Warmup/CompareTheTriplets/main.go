package main

import (
	"fmt"
)

func main() {
	a := [3]int{}
	b := [3]int{}
	var aSum int
	var bSum int
	for i := 0; i < 3; i++ {
		fmt.Scan(&a[i])
	}
	for i := 0; i < 3; i++ {
		fmt.Scan(&b[i])
	}
	for i := 0; i < 3; i++ {
		if a[i] > b[i] {
			aSum++
		} else if a[i] < b[i] {
			bSum++
		}
	}

	fmt.Printf("%d %d\n", aSum, bSum)
}
