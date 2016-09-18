package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	bio := bufio.NewReader(os.Stdin)
	line, _, err := bio.ReadLine()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hello, World.")
		fmt.Println(string(line))
	}

}
