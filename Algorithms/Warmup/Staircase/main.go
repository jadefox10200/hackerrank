package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	buildHashes(n)
}

func buildHashes(height int) {
	for i := 0; i < height; i++ {
		hashes := i + 1
		whiteSpace := height - hashes + 1
		for whiteSpace > 1 {
			fmt.Printf(" ")
			whiteSpace--
		}
		for hashes > 0 {
			fmt.Printf("#")
			hashes--
		}
		fmt.Printf("\n")
	}
}

/* more efficient alternate:
package main

import "fmt"
import "strconv"

func main() {
	var n int
	fmt.Scan(&n)
	str := ""
	for i := n; i > 0; i-- {
		str += "#"
		fmt.Printf("%"+strconv.Itoa(n)+"s\n", str)
	}
}
*/
