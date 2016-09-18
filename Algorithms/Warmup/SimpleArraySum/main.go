package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	sum := set(n)
	fmt.Println(sum)
}

func set(n int) int {
	a := make([]int, n)
	var sum int
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for y := 0; y < n; y++ {
		sum += a[y]
	}
	return sum
}

//PROBLEM:
// Given an array of  integers, can you find the sum of its elements?
//
// Input Format
//
// The first line contains an integer, , denoting the size of the array.
// The second line contains  space-separated integers representing the array's elements.
//
// Output Format
//
// Print the sum of the array's elements as a single integer.
//
// Sample Input
//
// 6
// 1 2 3 4 10 11
// Sample Output
//
// 31
// Explanation
//
// We print the sum of the array's elements, which is: .
