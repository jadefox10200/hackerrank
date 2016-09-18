package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Pull in the number of ints but don't need it:
	var testNumber int
	fmt.Scanf("%d", &testNumber)
	//Read in all of the ints and calculate one by one:
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		strNumber := scanner.Text()
		nLen := len(strNumber)
		var sum int
		for i := 0; i < nLen; i++ {
			//asserting:
			mainN, err := strconv.Atoi(strNumber)
			if err != nil {
				fmt.Println("Couldn't convert strNumber", err)
			}
			if strNumber[i] < 49 {
				continue
			}
			n := int(strNumber[i]) - 48
			if mainN%n == 0 {
				sum++
			}
		}
		fmt.Println(sum)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
