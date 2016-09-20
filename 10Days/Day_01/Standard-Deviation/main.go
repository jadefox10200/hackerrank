package main
import "fmt"
import "math"

func main() {
    var n int
    fmt.Scan(&n)
    var a = make([]float64, n)
    var r = make([]float64, n)
    for i := 0; i < n; i++{
        fmt.Scan(&a[i])
    }
    myMean := mean(n, a)
    for i := 0; i < n; i++ {
        r[i] = (a[i]-myMean) * (a[i]-myMean)
    }
    myMean = mean(n, r)
    fmt.Printf("%.1f", math.Sqrt(myMean))    
}

func mean(n int, a []float64) float64 {
	var sum float64
	for i := 0; i < n; i++ {
		sum += a[i]
	}
	sum = (sum / float64(n))
	return sum
}
