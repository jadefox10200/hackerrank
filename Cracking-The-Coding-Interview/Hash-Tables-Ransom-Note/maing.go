package main

import "fmt"

func main() {

	var textLen, noteLen int
	fmt.Scanf("%d %d", &textLen, &noteLen)
	textMap := make(map[string]int)
	noteMap := make(map[string]int)

	for i := 0; i < textLen; i++ {
		var word string
		fmt.Scan(&word)
		textMap[word]++

	}
	for i := 0; i < noteLen; i++ {
		var word string
		fmt.Scan(&word)
		noteMap[word]++
	}
	for k, _ := range noteMap {
		var check int
		check = textMap[k] - noteMap[k]
		if check < 0 {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")

}
