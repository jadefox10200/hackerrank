package main

import "fmt"

func main() {
	var s1, s2 string
	fmt.Scanln(&s1)
	fmt.Scanln(&s2)
	m1 := make(map[byte]int)
	m2 := make(map[byte]int)
	l1 := len(s1)
	l2 := len(s2)
	var counter int

	for i := 0; i < l1; i++ {
		m1[s1[i]-97]++
	}
	for i := 0; i < l2; i++ {
		m2[s2[i]-97]++
	}

	for i := 0; i < 26; i++ {
		counter += abs(m1[byte(i)] - m2[byte(i)])
	}

	fmt.Println(counter)

}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
