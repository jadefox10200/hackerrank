package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	c := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	c[0] = 1
	for i := 1; i < n; i++ {
		if a[i] > a[i-1] {
			c[i] = c[i-1] + 1
		} else {
			c[i] = 1
		}
	}
	for i := n - 2; i >= 0; i-- {
		if a[i] > a[i+1] {
			c[i] = max(c[i], c[i+1]+1)
		}
	}
	var sum int
	for _, v := range c {
		sum += v
	}
	fmt.Println(sum)
}

func max(x, y int) int {

	if x > y {
		return x
	}
	return y
}
