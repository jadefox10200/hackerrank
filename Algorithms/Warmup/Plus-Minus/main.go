package main

import "fmt"

func main() {
	var n float64
	fmt.Scan(&n)
	var val float64
	var pos, neg, zer float64
	for i := 0; i < int(n); i++ {
		fmt.Scan(&val)
		switch {
		case val > 0:
			pos++
		case val == 0:
			zer++
		case val < 0:
			neg++
		}
	}
	fmt.Println(checkValue(pos, n))
	fmt.Println(checkValue(neg, n))
	fmt.Println(checkValue(zer, n))
}

func checkValue(v, n float64) float64 {
	if v == 0 {
		return 0.0
	}
	return v / n
}
