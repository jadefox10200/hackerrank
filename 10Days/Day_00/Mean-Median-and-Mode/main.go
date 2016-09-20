package main

import "fmt"

func main() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	var n int
	fmt.Scan(&n)
	var array = make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}
	sort(n, array)
	theMean := mean(n, array)
	theMedian := median(n, array)
	theMode := mode(n, array)

	fmt.Printf("%.1f\n", theMean)
	fmt.Printf("%.1f\n", theMedian)
	fmt.Println(theMode)

}

func mean(n int, a []float64) float64 {
	var sum float64
	for i := 0; i < n; i++ {
		sum += a[i]
	}
	sum = (sum / float64(n))
	return sum
}

func median(n int, a []float64) float64 {
	if n%2 != 0 {
		middle := n / 2
		return a[middle-1]
	} else {
		m1 := n / 2
		m2 := m1 - 1
		return (a[m1] + a[m2]) / 2
	}
}

func mode(n int, a []float64) int {
	var result float64
	var TempCounter float64
	var ResultCounter float64
	result = a[0]
	for i := 0; i < n-1; i++ {
		if a[i] == a[i+1] {
			TempCounter++
			if TempCounter > ResultCounter {
				ResultCounter = TempCounter
				result = a[i]
			}
		} else {
			TempCounter = 0
		}
	}
	return int(result)
}

func sort(n int, a []float64) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
