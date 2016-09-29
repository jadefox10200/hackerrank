package main

import "fmt"

func main() {
	var l, n int
	fmt.Scanf("%d %d", &l, &n)
	var a = make([]int, l)

	for i := 0; i < l; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < l; i++ {
		fmt.Printf("%d ", a[(i+n)%l])
	}

}
