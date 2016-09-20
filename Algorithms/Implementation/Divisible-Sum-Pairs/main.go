package main
import "fmt"

func main() {
    
    var n int
    var k int
    var counter int
    fmt.Scanf("%d %d",&n, &k)
    var a = make([]int, n)
    for i := 0; i<n; i++ {
        fmt.Scan(&a[i])
    }
    for i := 0; i < n ; i++ {
        for j := i+1; j < n ; j++ {
            if (a[i]+a[j])%k == 0 {
                counter++
            }
        }
    }
    fmt.Println(counter)
    
    
}
