package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var value int
	var aLen int
	fmt.Scan(&value)
	fmt.Scan(&aLen)
	array := make([]int, aLen)
	for i := 0; i < aLen; i++ {
		fmt.Scan(&array[i])
		if array[i] == value {
			fmt.Println(i)
			break
		}
	}
}
