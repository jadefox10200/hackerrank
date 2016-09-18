package main

import (
    "fmt"
    "os"
    "bufio"
    "math"
)

func main() {	
    scanner := bufio.NewReader(os.Stdin)
    
    var mealCost float64
    var tipPercent float64
    var taxPercent float64   
    var totalCost float64
  
    fmt.Fscanln(scanner, &mealCost)
    fmt.Fscanln(scanner, &tipPercent)
    fmt.Fscanln(scanner, &taxPercent)
        
    tip := mealCost * tipPercent/100
    tax := mealCost * taxPercent/100

    totalCost = mealCost + tip + tax 

    result := Round(totalCost)
    fmt.Printf("The total meal cost is %d dollars.\n", int(result))
	
}

func Round(f float64) float64 {
    return math.Floor(f + .5)
}
