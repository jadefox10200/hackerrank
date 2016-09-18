package main
 import "fmt"
 func main() {
 //Enter your code here. Read input from STDIN. Print output to STDOUT
    var n int
    fmt.Scan(&n)
    if n % 2 != 0 {
        fmt.Println("Weird")
    }else if n <= 5{
        fmt.Println("Not Weird")
    } else if n <= 20 {
        fmt.Println("Weird")
    } else {
        fmt.Println("Not Weird")
    }
}
