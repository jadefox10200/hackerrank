package main

import "fmt"

func main() {
	var n, k, q, j, m int
	fmt.Scanf("%d %d %d", &n, &k, &q)
	k %= n
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	for i := 0; i < q; i++ {
		fmt.Scan(&m)
		j = m - k
		if j < 0 {
			fmt.Println(a[n+j])
		} else {
			fmt.Println(a[j])
		}
	}
}
