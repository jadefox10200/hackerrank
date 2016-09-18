package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var i uint32 = 4
	var d float32 = 4.0
	var s string = "HackerRank "

	scanner := bufio.NewReader(os.Stdin)

	var i2 uint32
	var d2 float32
	var s2 string

	fmt.Fscanln(scanner, &i2)
	fmt.Fscanln(scanner, &d2)
	myScanner := bufio.NewScanner(scanner)

	var dResult = d + d2
	fmt.Println(i + i2)
	fmt.Printf("%.1f\n", dResult)
	for myScanner.Scan() {
		s2 = myScanner.Text()
	}
	fmt.Println(s + s2)
}

// Declare second integer, double, and String variables.

// Read and save an integer, double, and String to your variables.

// Print the sum of both integer variables on a new line.

// Print the sum of the double variables on a new line.

// Concatenate and print the String variables on a new line
// The 's' variable above should be printed first.
