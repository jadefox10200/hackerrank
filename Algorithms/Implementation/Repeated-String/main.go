package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var s string
	fmt.Scan(&s)
	var n int
	fmt.Scan(&n)
	var a int
	var l = len(s)
	var rem = n % l

	a = getCount(s)
	a = a * (n / l)
	if (n % l) == 0 {
		fmt.Printf("%d", a)
	} else {
		fmt.Printf("%d", a+getCount(s[:rem]))
	}
}

func getCount(s string) int {
	var a int
	var l = len(s)
	for i := 0; i < l; i++ {
		if s[i] == 'a' {
			a++
		}
	}
	return a
}
