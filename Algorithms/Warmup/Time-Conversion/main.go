package main

import "fmt"

func main() {

	var h, m, s int
	var ap string
	fmt.Scanf("%d:%d:%d%s", &h, &m, &s, &ap)
	if ap == "PM" && h != 12 {
		h += 12
	}
	if ap == "AM" && h == 12 {
		h = 0
	}
	fmt.Printf("%02d:%02d:%02d", h, m, s)
}
