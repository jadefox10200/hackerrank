package main

import "fmt"

func main() {
	var max = -100
	var sum = -100
	var a = [6][6]int{}
	for i := 0; i < 6; i++ {
		for k := 0; k < 6; k++ {
			fmt.Scan(&a[i][k])
		}
	}

	for i := 0; i < 4; i++ {
		for k := 0; k < 4; k++ {
			sum = a[i][k] + a[i][k+1] + a[i][k+2] + a[i+1][k+1] + a[i+2][k] + a[i+2][k+1] + a[i+2][k+2]
			if sum > max {
				max = sum
			}
		}
	}
	fmt.Println(max)
}
