package main
import "fmt"

func main() {
    var n int
    fmt.Scan(&n)
    var e = make([]float64, n)
    var f = make([]float64, n)
    var s []float64
    
    for i := 0; i < n; i++ {
        fmt.Scan(&e[i])
    }
    for i := 0; i < n; i++ {
        fmt.Scan(&f[i])
    }
    for i := 0; i< n;i++ {
        for j := 0; j < int(f[i]); j++{
            s = append(s, e[i])
        }
    }
    l := len(s)
    sort(l,s)
    q1,_,q3 := quartiles(l,s)
    fmt.Printf("%.1f\n", q3-q1)
    
    
}

func quartiles(n int, a []float64) (float64, float64, float64){
    var q1,q2,q3 float64
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

func median(a []float64) float64 {
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

func sort(n int, a []float64) {
	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
