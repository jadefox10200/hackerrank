package main

import "fmt"

func main() {
	var size int
	var sharedIterator int
	var sharedValue int
	var buff int
	var charged int

	fmt.Scanf("%d %d\n", &size, &sharedIterator)

	for i := 0; i < size; i++ {
		fmt.Scan(&buff)
		if i == sharedIterator {
			continue
		}
		sharedValue += buff
	}

	fmt.Scanf("%d", &charged)
	if (sharedValue / 2) == charged {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(charged - (sharedValue / 2))
	}
}
