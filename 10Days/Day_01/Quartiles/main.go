package main
import "fmt"
//import "math"

func main() {
    var n int 
    fmt.Scan(&n)
    var a = make([]int, n)
    for i := 0; i < n; i++ {
        fmt.Scan(&a[i])
    }
    sort(n, a)
    q1, q2, q3 := quartiles(n, a)    
    fmt.Println(q1)
    fmt.Println(q2)
    fmt.Println(q3)
}

func quartiles(n int, a []int) (int, int, int){
    var q1,q2,q3 int
    q2 = median(a)
    if n%2 != 0 {
		middle := n / 2
        left := a[:middle]
        right := a[middle+1:]
        q1 = median(left)
        q3 = median(right)
		return q1,q2,q3
	} else {
		middle := n / 2
        left := a[:middle]
        right := a[middle:]
        q1 = median(left)
        q3 = median(right)
		return q1,q2,q3
	}
}


func median(a []int) (int) {
    n := len(a)
	if n%2 != 0 {
		middle := n / 2
		return a[middle]
	} else {
		m1 := n / 2
		m2 := m1 - 1
		return (a[m1] + a[m2]) / 2
	}
}

func sort(n int, a []int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
