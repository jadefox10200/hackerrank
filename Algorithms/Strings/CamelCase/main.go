package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var counter = 1
	buf := bufio.NewReader(os.Stdin)
	s := bufio.NewScanner(buf)
	s.Split(bufio.ScanRunes)
	for i := 0; s.Scan(); i++ {
		b := s.Bytes()
		if b[0] < 97 {
			counter++
		}

	}
	fmt.Println(counter)
}
