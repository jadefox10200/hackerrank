package main

import "fmt"

func main() {
	var n, left, right int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var value int
			fmt.Scan(&value)
			if i == j {
				left += value
			}
			if i+j == n-1 {
				right += value
			}
		}
	}
	fmt.Println(abs(left - right))

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	if x == 0 {
		return 0
	}
	return x
}
