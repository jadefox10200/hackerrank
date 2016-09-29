package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var students, acceptance int
		fmt.Scanf("%d %d\n", &students, &acceptance)
		var onTimeCount int
		var currStudent int
		for j := 0; j < students; j++ {
			fmt.Scan(&currStudent)
			if currStudent <= 0 {
				onTimeCount++
			}
		}
		if onTimeCount >= acceptance {
			fmt.Println("NO")
		} else {
			fmt.Println("YES")
		}

	}
}
