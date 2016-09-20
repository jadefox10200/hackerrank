package main
import "fmt"

func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
  var n int
  fmt.Scan(&n)
  for i := 1; i <= 10; i++ {
      result := n * i
      fmt.Printf("%d x %d = %d\n", n,i,result)
  }
}
