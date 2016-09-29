package main

import "fmt"

func main() {
	var k1, v1, k2, v2 int
	fmt.Scanf("%d %d %d %d", &k1, &v1, &k2, &v2)
	if (k1 < k2) && (v1 < v2) {
		fmt.Println("NO")
	} else {
		if (v1 != v2) && ((k2-k1)%(v1-v2)) == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
