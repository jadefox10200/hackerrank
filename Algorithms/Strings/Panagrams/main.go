package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var panagram = "pangram"
	alpha := [26]int{}
	buf := bufio.NewReader(os.Stdin)
	s := bufio.NewScanner(buf)
	s.Split(bufio.ScanRunes)
	for i := 0; s.Scan(); i++ {
		b := s.Bytes()
		if b[0] == 32 || b[0] > 122 {
			continue
		} else if b[0] < 97 {
			b[0] += 32
		}
		n := b[0] - 96
		alpha[n-1] = int(n)
	}
	for i := 0; i < 26; i++ {
		if alpha[i] == 0 {
			panagram = "not pangram"
		}
	}
	fmt.Println(panagram)
}
