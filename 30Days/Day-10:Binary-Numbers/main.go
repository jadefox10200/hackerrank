package main

import "fmt"
import "strconv"

func main() {
	var n int64
	fmt.Scan(&n)
	si := strconv.FormatInt(n, 2)
	var counter = 1
	var max = 1
	l := len(si)
	for i := 0; i < l-1; i++ {
		if int(si[i]-48) == 1 && int(si[i+1]-48) == 1 {
			counter++
			if counter > max {
				max = counter
			}
		} else {
			counter = 1
		}
	}
	fmt.Println(max)
}
